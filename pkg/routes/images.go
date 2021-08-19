package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/redhatinsights/edge-api/config"
	"github.com/redhatinsights/edge-api/pkg/clients/imagebuilder"
	"github.com/redhatinsights/edge-api/pkg/db"
	"github.com/redhatinsights/edge-api/pkg/errors"
	"github.com/redhatinsights/edge-api/pkg/models"
	"github.com/redhatinsights/edge-api/pkg/routes/common"
	"github.com/redhatinsights/edge-api/pkg/services"
	log "github.com/sirupsen/logrus"
)

// This provides type safety in the context object for our "image" key.  We
// _could_ use a string but we shouldn't just in case someone else decides that
// "image" would make the perfect key in the context object.  See the
// documentation: https://golang.org/pkg/context/#WithValue for further
// rationale.
type imageTypeKey int

const imageKey imageTypeKey = iota

// MakeImageRouter adds support for operations on images
func MakeImagesRouter(sub chi.Router) {
	sub.With(validateGetAllImagesSearchParams).With(common.Paginate).Get("/", GetAllImages)
	sub.Get("/reserved-usernames", GetReservedUsernames)
	sub.Post("/", CreateImage)
	sub.Route("/{imageId}", func(r chi.Router) {
		r.Use(ImageCtx)
		r.Get("/", GetImageByID)
		r.Get("/status", GetImageStatusByID)
		r.Get("/repo", GetRepoForImage)
		r.Get("/metadata", GetMetadataForImage)
		r.Post("/installer", CreateInstallerForImage)
		r.Post("/repo", CreateRepoForImage)
		r.Post("/kickstart", CreateKickStartForImage)
	})
}

var validStatuses = []string{models.ImageStatusCreated, models.ImageStatusBuilding, models.ImageStatusError, models.ImageStatusSuccess}

// WaitGroup is the waitg roup for pending image builds
var WaitGroup sync.WaitGroup

// ImageCtx is a handler for Image requests
func ImageCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var image models.Image
		account, err := common.GetAccount(r)
		if err != nil {
			err := errors.NewBadRequest(err.Error())
			w.WriteHeader(err.Status)
			json.NewEncoder(w).Encode(&err)
			return
		}
		if imageID := chi.URLParam(r, "imageId"); imageID != "" {
			id, err := strconv.Atoi(imageID)
			if err != nil {
				err := errors.NewBadRequest(err.Error())
				w.WriteHeader(err.Status)
				json.NewEncoder(w).Encode(&err)
				return
			}
			result := db.DB.Where("images.account = ?", account).Joins("Commit").First(&image, id)
			if image.InstallerID != nil {
				result := db.DB.First(&image.Installer, image.InstallerID)
				if result.Error != nil {
					err := errors.NewInternalServerError()
					w.WriteHeader(err.Status)
					json.NewEncoder(w).Encode(&err)
					return
				}
			}
			err = db.DB.Model(&image.Commit).Association("Packages").Find(&image.Commit.Packages)
			if err != nil {
				err := errors.NewInternalServerError()
				w.WriteHeader(err.Status)
				json.NewEncoder(w).Encode(&err)
				return
			}
			if result.Error != nil {
				err := errors.NewNotFound(result.Error.Error())
				w.WriteHeader(err.Status)
				json.NewEncoder(w).Encode(&err)
				return
			}
			ctx := context.WithValue(r.Context(), imageKey, &image)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}

// A CreateImageRequest model.
type CreateImageRequest struct {
	// The image to create.
	//
	// in: body
	// required: true
	Image *models.Image
}

func createImage(image *models.Image, account string, ctx context.Context) error {
	client := imagebuilder.InitClient(ctx)
	image, err := client.ComposeCommit(image)
	if err != nil {
		return err
	}
	image.Account = account
	image.Commit.Account = account
	image.Commit.Status = models.ImageStatusBuilding
	image.Status = models.ImageStatusBuilding
	if image.ImageType == models.ImageTypeInstaller {
		image.Installer.Status = models.ImageStatusCreated
		image.Installer.Account = image.Account
		tx := db.DB.Create(&image.Installer)
		if tx.Error != nil {
			return tx.Error
		}
	}
	tx := db.DB.Create(&image.Commit)
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.DB.Create(&image)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func createRepoForImage(i *models.Image, ctx context.Context) *models.Repo {
	log.Infof("Commit %d for Image %d is ready. Creating OSTree repo.", i.Commit.ID, i.ID)
	repo := &models.Repo{
		CommitID: i.Commit.ID,
		Commit:   i.Commit,
		Status:   models.RepoStatusBuilding,
	}
	tx := db.DB.Create(repo)
	if tx.Error != nil {
		log.Error(tx.Error)
		panic(tx.Error)
	}
	rb := services.NewRepoBuilder(ctx)
	repo, err := rb.ImportRepo(repo)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	log.Infof("OSTree repo %d for commit %d and Image %d is ready. ", repo.ID, i.Commit.ID, i.ID)

	return repo
}

func setErrorStatusOnImage(err error, i *models.Image) {
	i.Status = models.ImageStatusError
	tx := db.DB.Save(i)
	if tx.Error != nil {
		panic(tx.Error)
	}
	if i.Commit != nil {
		i.Commit.Status = models.ImageStatusError
		tx := db.DB.Save(i.Commit)
		if tx.Error != nil {
			panic(tx.Error)
		}
	}
	if i.Installer != nil {
		i.Installer.Status = models.ImageStatusError
		tx := db.DB.Save(i.Installer)
		if tx.Error != nil {
			panic(tx.Error)
		}
	}
	if err != nil {
		log.Error(err)
		panic(err)
	}
}

func postProcessImage(id uint, ctx context.Context) {
	WaitGroup.Add(1) // Processing one image

	defer func() {
		WaitGroup.Done() // Done with one image (sucessfuly or not)
		fmt.Println("Shutting down go routine")
		if err := recover(); err != nil {
			log.Fatalf("%s", err)
		}
	}()
	var i *models.Image
	db.DB.Joins("Commit").Joins("Installer").First(&i, id)

	client := imagebuilder.InitClient(ctx)

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		sig := <-sigint
		// Reload image to get updated status
		db.DB.Joins("Commit").Joins("Installer").First(&i, i.ID)
		if i.Status == models.ImageStatusBuilding {
			log.Infof("Captured %v, marking image as error", sig)
			setErrorStatusOnImage(nil, i)
			WaitGroup.Done()
		}
	}()
	for {
		i, err := updateImageStatus(i, ctx)
		if err != nil {
			setErrorStatusOnImage(err, i)
		}
		if i.Commit.Status != models.ImageStatusBuilding {
			break
		}
		time.Sleep(1 * time.Minute)
	}

	go func() {
		i, err := client.GetMetadata(i)
		if err != nil {
			log.Error(err)
		} else {
			db.DB.Save(&i.Commit)
		}
	}()

	repo := createRepoForImage(i, ctx)

	// TODO: We need to discuss this whole thing post-July deliverable
	if i.ImageType == models.ImageTypeInstaller {
		i, err := client.ComposeInstaller(repo, i)
		if err != nil {
			log.Error(err)
			setErrorStatusOnImage(err, i)
		}
		i.Installer.Status = models.ImageStatusBuilding
		tx := db.DB.Save(&i.Installer)
		if tx.Error != nil {
			log.Error(err)
			setErrorStatusOnImage(err, i)
		}

		for {
			i, err := updateImageStatus(i, ctx)
			if err != nil {
				setErrorStatusOnImage(err, i)
			}
			if i.Installer.Status != models.ImageStatusBuilding {
				break
			}
			time.Sleep(1 * time.Minute)
		}

		if i.Installer.Status == models.ImageStatusSuccess {
			err = addUserInfo(i)
			if err != nil {
				// TODO: Temporary. Handle error better.
				log.Errorf("Kickstart file injection failed %s", err.Error())
			}
		}
	}

	log.Infof("Setting image %d status as success", i.ID)
	if i.Commit.Status == models.ImageStatusSuccess {
		if i.Installer != nil || i.Installer.Status == models.ImageStatusSuccess {
			i.Status = models.ImageStatusSuccess
			db.DB.Save(&i)
		}
	}
}

// CreateImage creates an image on hosted image builder.
// It always creates a commit on Image Builder.
// We're creating a update on the background to transfer the commit to our repo.
func CreateImage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var image *models.Image
	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
		log.Error(err)
		err := errors.NewInternalServerError()
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(&err)
		return
	}
	if err := image.ValidateRequest(); err != nil {
		log.Info(err)
		err := errors.NewBadRequest(err.Error())
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(&err)
		return
	}
	account, err := common.GetAccount(r)
	if err != nil {
		log.Info(err)
		err := errors.NewBadRequest(err.Error())
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(&err)
		return
	}
	err = createImage(image, account, r.Context())
	if err != nil {
		log.Error(err)
		err := errors.NewInternalServerError()
		err.Title = "Failed creating image"
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(&err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&image)

	go postProcessImage(image.ID, r.Context())
}

var imageFilters = common.ComposeFilters(
	common.OneOfFilterHandler(&common.Filter{
		QueryParam: "status",
		DBField:    "images.status",
	}),
	common.ContainFilterHandler(&common.Filter{
		QueryParam: "name",
		DBField:    "images.name",
	}),
	common.ContainFilterHandler(&common.Filter{
		QueryParam: "distribution",
		DBField:    "images.distribution",
	}),
	common.CreatedAtFilterHandler(&common.Filter{
		QueryParam: "created_at",
		DBField:    "images.created_at",
	}),
	common.SortFilterHandler("images", "created_at", "DESC"),
)

type validationError struct {
	Key    string
	Reason string
}

func validateGetAllImagesSearchParams(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errs := []validationError{}
		if statuses, ok := r.URL.Query()["status"]; ok {
			for _, status := range statuses {
				if status != models.ImageStatusCreated && status != models.ImageStatusBuilding && status != models.ImageStatusError && status != models.ImageStatusSuccess {
					errs = append(errs, validationError{Key: "status", Reason: fmt.Sprintf("%s is not a valid status. Status must be %s", status, strings.Join(validStatuses, " or "))})
				}
			}
		}
		if val := r.URL.Query().Get("created_at"); val != "" {
			if _, err := time.Parse(common.LayoutISO, val); err != nil {
				errs = append(errs, validationError{Key: "created_at", Reason: err.Error()})
			}
		}
		if val := r.URL.Query().Get("sort_by"); val != "" {
			name := val
			if string(val[0]) == "-" {
				name = val[1:]
			}
			if name != "status" && name != "name" && name != "distribution" && name != "created_at" {
				errs = append(errs, validationError{Key: "sort_by", Reason: fmt.Sprintf("%s is not a valid sort_by. Sort-by must be status or name or distribution or created_at", name)})
			}
		}

		if len(errs) == 0 {
			next.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&errs)
	})
}

// GetAllImages image objects from the database for an account
func GetAllImages(w http.ResponseWriter, r *http.Request) {
	var count int64
	var images []models.Image
	result := imageFilters(r, db.DB)
	pagination := common.GetPagination(r)
	account, err := common.GetAccount(r)
	if err != nil {
		log.Info(err)
		err := errors.NewBadRequest(err.Error())
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(&err)
		return
	}
	countResult := imageFilters(r, db.DB.Model(&models.Image{})).Where("images.account = ?", account).Count(&count)
	if countResult.Error != nil {
		countErr := errors.NewInternalServerError()
		log.Error(countErr)
		w.WriteHeader(countErr.Status)
		json.NewEncoder(w).Encode(&countErr)
		return
	}
	result = result.Limit(pagination.Limit).Offset(pagination.Offset).Where("images.account = ?", account).Joins("Commit").Joins("Installer").Find(&images)
	if result.Error != nil {
		log.Error(err)
		err := errors.NewInternalServerError()
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(&err)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"data": &images, "count": count})
}

func getImage(w http.ResponseWriter, r *http.Request) *models.Image {
	ctx := r.Context()
	image, ok := ctx.Value(imageKey).(*models.Image)
	if !ok {
		err := errors.NewBadRequest("Must pass image id")
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(&err)
		return nil
	}
	return image
}

func updateImageStatus(image *models.Image, ctx context.Context) (*models.Image, error) {
	client := imagebuilder.InitClient(ctx)
	if image.Commit.Status == models.ImageStatusBuilding {
		image, err := client.GetCommitStatus(image)
		if err != nil {
			return image, err
		}
		if image.Commit.Status != models.ImageStatusBuilding {
			tx := db.DB.Save(&image.Commit)
			if tx.Error != nil {
				return image, tx.Error
			}
		}
	}
	if image.Installer != nil && image.Installer.Status == models.ImageStatusBuilding {
		image, err := client.GetInstallerStatus(image)
		if err != nil {
			return image, err
		}
		if image.Installer.Status != models.ImageStatusBuilding {
			tx := db.DB.Save(&image.Installer)
			if tx.Error != nil {
				return image, tx.Error
			}
		}
	}
	if image.Status != models.ImageStatusBuilding {
		tx := db.DB.Save(&image)
		if tx.Error != nil {
			return image, tx.Error
		}
	}
	return image, nil
}

// GetImageStatusByID returns the image status.
func GetImageStatusByID(w http.ResponseWriter, r *http.Request) {
	if image := getImage(w, r); image != nil {
		json.NewEncoder(w).Encode(struct {
			Status string
			Name   string
			ID     uint
		}{
			image.Status,
			image.Name,
			image.ID,
		})
	}
}

// GetImageByID obtains a image from the database for an account
func GetImageByID(w http.ResponseWriter, r *http.Request) {
	if image := getImage(w, r); image != nil {
		json.NewEncoder(w).Encode(image)
	}
}

// CreateInstallerForImage creates a installer for a Image
// It requires a created image and an update for the commit
func CreateInstallerForImage(w http.ResponseWriter, r *http.Request) {
	image := getImage(w, r)
	var imageInstaller *models.Installer
	if err := json.NewDecoder(r.Body).Decode(&imageInstaller); err != nil {
		log.Error(err)
		err := errors.NewInternalServerError()
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(&err)
		return
	}
	image.ImageType = models.ImageTypeInstaller
	image.Installer = imageInstaller

	tx := db.DB.Save(&image)
	if tx.Error != nil {
		log.Error(tx.Error)
		err := errors.NewInternalServerError()
		err.Title = "Failed saving image status"
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(&err)
		return
	}
	repoService := services.NewRepoService()
	repo, err := repoService.GetRepoByCommitID(image.CommitID)
	if err != nil {
		err := errors.NewBadRequest(fmt.Sprintf("Commit Repo wasn't found in the database: #%v", image.Commit.ID))
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(&err)
		return
	}
	client := imagebuilder.InitClient(r.Context())
	image, err = client.ComposeInstaller(repo, image)
	if err != nil {
		log.Error(err)
		err := errors.NewInternalServerError()
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(&err)
		return
	}
	image.Installer.Status = models.ImageStatusBuilding
	image.Status = models.ImageStatusBuilding
	tx = db.DB.Save(&image)
	if tx.Error != nil {
		log.Error(err)
		err := errors.NewInternalServerError()
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(&err)
		return
	}
	tx = db.DB.Save(&image.Installer)
	if tx.Error != nil {
		log.Error(err)
		err := errors.NewInternalServerError()
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(&err)
		return
	}

	go func(id uint, ctx context.Context) {
		var i *models.Image
		db.DB.Joins("Commit").Joins("Installer").First(&i, id)
		for {
			i, err := updateImageStatus(i, ctx)
			if err != nil {
				setErrorStatusOnImage(err, i)
			}
			if i.Installer.Status != models.ImageStatusBuilding {
				break
			}
			time.Sleep(1 * time.Minute)
		}
		if i.Installer.Status == models.ImageStatusSuccess {
			err = addUserInfo(image)
			if err != nil {
				// TODO: Temporary. Handle error better.
				log.Errorf("Kickstart file injection failed %s", err.Error())
			}
		}
	}(image.ID, r.Context())

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&image)
}

// CreateRepoForImage creates a repo for a Image
func CreateRepoForImage(w http.ResponseWriter, r *http.Request) {
	image := getImage(w, r)

	go func(id uint) {
		var i *models.Image
		db.DB.Joins("Commit").Joins("Installer").First(&i, id)
		db.DB.First(&i.Commit, i.CommitID)
		createRepoForImage(i, r.Context())
	}(image.ID)

	w.WriteHeader(http.StatusOK)
}

// Download the ISO, inject the kickstart with username and ssh key
// re upload the ISO
func addUserInfo(image *models.Image) error {
	// Absolute path for manipulating ISO's
	destPath := "/var/tmp/"

	downloadURL := image.Installer.ImageBuildISOURL
	sshKey := image.Installer.SSHKey
	username := image.Installer.Username
	// Files that will be used to modify the ISO and will be cleaned
	imageName := destPath + image.Name
	kickstart := destPath + "finalKickstart-" + username + ".ks"

	err := downloadISO(imageName, downloadURL)
	if err != nil {
		return fmt.Errorf("error downloading ISO file :: %s", err.Error())
	}

	err = addSSHKeyToKickstart(sshKey, username, kickstart)
	if err != nil {
		return fmt.Errorf("error adding ssh key to kickstart file :: %s", err.Error())
	}

	err = exeInjectionScript(kickstart, imageName, image.ID)
	if err != nil {
		return fmt.Errorf("error execuiting fleetkick script :: %s", err.Error())
	}

	err = uploadISO(image, imageName)
	if err != nil {
		return fmt.Errorf("error uploading ISO :: %s", err.Error())
	}

	err = cleanFiles(kickstart, imageName, image.ID)
	if err != nil {
		return fmt.Errorf("error cleaning files :: %s", err.Error())
	}

	return nil
}

// UnameSSH is the template struct for username and ssh key
type UnameSSH struct {
	Sshkey   string
	Username string
}

// Adds user provided ssh key to the kickstart file.
func addSSHKeyToKickstart(sshKey string, username string, kickstart string) error {
	cfg := config.Get()

	td := UnameSSH{sshKey, username}

	log.Infof("Opening file %s", cfg.TemplatesPath)
	t, err := template.ParseFiles(cfg.TemplatesPath + "templateKickstart.ks")
	if err != nil {
		return err
	}

	log.Infof("Creating file %s", kickstart)
	file, err := os.Create(kickstart)
	if err != nil {
		return err
	}

	log.Infof("Injecting username %s and key %s into template", username, sshKey)
	err = t.Execute(file, td)
	if err != nil {
		return err
	}
	file.Close()

	return nil
}

// Download created ISO into the file system.
func downloadISO(isoName string, url string) error {
	log.Infof("Creating iso %s", isoName)
	iso, err := os.Create(isoName)
	if err != nil {
		return err
	}
	defer iso.Close()

	log.Infof("Downloading ISO %s", url)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.Copy(iso, res.Body)
	if err != nil {
		return err
	}

	return nil
}

// Upload finished ISO to S3
func uploadISO(image *models.Image, imageName string) error {

	uploadPath := fmt.Sprintf("%s/isos/%s.iso", image.Account, image.Name)
	filesService := services.NewFilesService()
	url, err := filesService.Uploader.UploadFile(imageName, uploadPath)

	if err != nil {
		return fmt.Errorf("error uploading the ISO :: %s :: %s", uploadPath, err.Error())
	}

	image.Installer.ImageBuildISOURL = url
	tx := db.DB.Save(&image.Installer)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Remove edited kickstart after use.
func cleanFiles(kickstart string, isoName string, imageID uint) error {
	err := os.Remove(kickstart)
	if err != nil {
		return err
	}
	log.Info("Kickstart file " + kickstart + " removed!")

	err = os.Remove(isoName)
	if err != nil {
		return err
	}
	log.Info("ISO file " + isoName + " removed!")

	workDir := fmt.Sprintf("/var/tmp/workdir%d", imageID)
	err = os.RemoveAll(workDir)
	if err != nil {
		return err
	}
	log.Info("work dir file " + workDir + " removed!")

	return nil
}

//GetRepoForImage gets the repository for a Image
func GetRepoForImage(w http.ResponseWriter, r *http.Request) {
	if image := getImage(w, r); image != nil {
		repoService := services.NewRepoService()
		repo, err := repoService.GetRepoByCommitID(image.CommitID)
		if err != nil {
			err := errors.NewNotFound(fmt.Sprintf("Commit repo wasn't found in the database: #%v", image.Commit.ID))
			w.WriteHeader(err.Status)
			json.NewEncoder(w).Encode(&err)
			return
		}
		json.NewEncoder(w).Encode(repo)
	}
}

//GetMetadataForImage gets the metadata from image-builder on /metadata endpoint
func GetMetadataForImage(w http.ResponseWriter, r *http.Request) {
	client := imagebuilder.InitClient(r.Context())
	if image := getImage(w, r); image != nil {
		meta, err := client.GetMetadata(image)
		if err != nil {
			log.Fatal(err)
		}
		if image.Commit.OSTreeCommit != "" {
			tx := db.DB.Save(&image.Commit)
			if tx.Error != nil {
				panic(tx.Error)
			}
		}
		json.NewEncoder(w).Encode(meta)
	}
}

// CreateKickStartForImage creates a kickstart file for an existent image
func CreateKickStartForImage(w http.ResponseWriter, r *http.Request) {
	if image := getImage(w, r); image != nil {
		err := addUserInfo(image)
		if err != nil {
			// TODO: Temporary. Handle error better.
			log.Errorf("Kickstart file injection failed %s", err.Error())
			err := errors.NewInternalServerError()
			w.WriteHeader(err.Status)
			json.NewEncoder(w).Encode(&err)
			return
		}
	}
}

// Inject the custom kickstart into the iso via mkksiso.
func exeInjectionScript(kickstart string, image string, imageID uint) error {
	fleetBashScript := "/usr/local/bin/fleetkick.sh"
	workDir := fmt.Sprintf("/var/tmp/workdir%d", imageID)
	err := os.Mkdir(workDir, 0755)
	if err != nil {
		return err
	}

	cmd := exec.Command(fleetBashScript, kickstart, image, image, workDir)
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	log.Infof("fleetkick output: %s\n", output)
	return nil
}

func GetReservedUsernames(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.ReservedUsernames)
}
