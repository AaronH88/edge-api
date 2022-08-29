package services

import (
	"context"

	"github.com/redhatinsights/edge-api/pkg/db"
	"github.com/redhatinsights/edge-api/pkg/models"
	"github.com/redhatinsights/edge-api/pkg/routes/common"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// CommitServiceInterface defines the interface to handle the business logic of RHEL for Edge Commits
type CommitServiceInterface interface {
	GetCommitByID(commitID uint, orgID string) (*models.Commit, error)
	GetCommitByOSTreeCommit(ost string) (*models.Commit, error)
	ValidateDevicesImageSetWithCommit(deviceUUID []string, commitID uint) error
}

// NewCommitService gives a instance of the main implementation of CommitServiceInterface
func NewCommitService(ctx context.Context, log *log.Entry) CommitServiceInterface {
	return &CommitService{
		Service: Service{ctx: ctx, log: log.WithField("service", "commit")},
	}
}

// CommitService is the main implementation of a CommitServiceInterface
type CommitService struct {
	Service
}

// GetCommitByID receives CommitID uint and get a *models.Commit back
func (s *CommitService) GetCommitByID(commitID uint, orgID string) (*models.Commit, error) {

	s.log = s.log.WithField("commitID", commitID)
	s.log.Debug("Getting commit by id")
	var commit models.Commit
	result := db.Org(orgID, "").First(&commit, commitID)
	if result.Error != nil {
		s.log.WithField("error", result.Error.Error()).Error("Error searching for commit by commitID")
		return nil, result.Error
	}
	s.log.Debug("Commit retrieved")
	return &commit, nil
}

// GetCommitByOSTreeCommit receives an OSTreeCommit string and get a *models.Commit back
func (s *CommitService) GetCommitByOSTreeCommit(ost string) (*models.Commit, error) {
	s.log = s.log.WithField("ostreeCommitHash", ost)
	var commit models.Commit
	result := db.DB.Where("os_tree_commit = ?", ost).First(&commit)
	if result.Error != nil {
		s.log.WithField("error", result.Error.Error()).Error("Error searching for commit by ostreeCommitHash")
		return nil, result.Error
	}
	s.log.Debug("Commit retrieved")
	return &commit, nil
}

// ValidateDevicesImageSetWithCommit validates if user provided commitID belong to same ImageSet as of Device Image
func (s *CommitService) ValidateDevicesImageSetWithCommit(devicesUUID []string, commitID uint) error {
	type ImageSetsDevices struct {
		ImageSetID   uint `json:"image_set_id"`
		DevicesCount int  `json:"devices_count"`
	}
	var imageSetsDevices []ImageSetsDevices
	var commitImage models.Image
	orgID, err := common.GetOrgIDFromContext(s.ctx)
	if err != nil {
		return err
	}

	if result := db.Org(orgID, "devices").Table("devices").
		Select(`images.image_set_id as "image_set_id", count(devices.id) as devices_count`).
		Joins("JOIN images ON devices.image_id = images.id").
		Where("devices.uuid in (?)", devicesUUID).
		Group("images.image_set_id").
		Find(&imageSetsDevices); result.Error != nil {
		s.log.WithField("error", result.Error.Error()).Error("Error searching for ImageSet of Device Images")
		return result.Error
	}

	if len(imageSetsDevices) == 0 {
		return new(ImageSetNotFoundError)
	}
	if len(imageSetsDevices) > 1 {
		return new(DeviceHasMoreThanOneImageSet)
	}
	imageSetDevices := imageSetsDevices[0]
	if imageSetDevices.DevicesCount != len(devicesUUID) {
		return new(ImageSetNotFoundForAllDevices)

	}

	if result := db.Org(orgID, "").Where("commit_id = ?", commitID).First(&commitImage); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return new(CommitImageNotFound)
		}
		s.log.WithField("error", result.Error.Error()).Error("Error searching for Images using user provided CommitID")
		return result.Error
	}

	if commitImage.ImageSetID == nil {
		return new(ImageHasNoImageSet)
	}
	if imageSetDevices.ImageSetID != *commitImage.ImageSetID {
		return new(InvalidCommitID)
	}
	return nil

}
