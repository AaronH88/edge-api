// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/images.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/redhatinsights/edge-api/pkg/models"
	services "github.com/redhatinsights/edge-api/pkg/services"
	reflect "reflect"
)

// MockImageServiceInterface is a mock of ImageServiceInterface interface
type MockImageServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockImageServiceInterfaceMockRecorder
}

// MockImageServiceInterfaceMockRecorder is the mock recorder for MockImageServiceInterface
type MockImageServiceInterfaceMockRecorder struct {
	mock *MockImageServiceInterface
}

// NewMockImageServiceInterface creates a new mock instance
func NewMockImageServiceInterface(ctrl *gomock.Controller) *MockImageServiceInterface {
	mock := &MockImageServiceInterface{ctrl: ctrl}
	mock.recorder = &MockImageServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImageServiceInterface) EXPECT() *MockImageServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateImage mocks base method
func (m *MockImageServiceInterface) CreateImage(image *models.Image, account string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateImage", image, account)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateImage indicates an expected call of CreateImage
func (mr *MockImageServiceInterfaceMockRecorder) CreateImage(image, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateImage", reflect.TypeOf((*MockImageServiceInterface)(nil).CreateImage), image, account)
}

// ResumeBuilds mocks base method
func (m *MockImageServiceInterface) ResumeBuilds() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResumeBuilds")
}

// ResumeBuilds indicates an expected call of ResumeBuilds
func (mr *MockImageServiceInterfaceMockRecorder) ResumeBuilds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResumeBuilds", reflect.TypeOf((*MockImageServiceInterface)(nil).ResumeBuilds))
}

// UpdateImage mocks base method
func (m *MockImageServiceInterface) UpdateImage(image, previousImage *models.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateImage", image, previousImage)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateImage indicates an expected call of UpdateImage
func (mr *MockImageServiceInterfaceMockRecorder) UpdateImage(image, previousImage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateImage", reflect.TypeOf((*MockImageServiceInterface)(nil).UpdateImage), image, previousImage)
}

// AddUserInfo mocks base method
func (m *MockImageServiceInterface) AddUserInfo(image *models.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserInfo", image)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserInfo indicates an expected call of AddUserInfo
func (mr *MockImageServiceInterfaceMockRecorder) AddUserInfo(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserInfo", reflect.TypeOf((*MockImageServiceInterface)(nil).AddUserInfo), image)
}

// UpdateImageStatus mocks base method
func (m *MockImageServiceInterface) UpdateImageStatus(image *models.Image) (*models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateImageStatus", image)
	ret0, _ := ret[0].(*models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateImageStatus indicates an expected call of UpdateImageStatus
func (mr *MockImageServiceInterfaceMockRecorder) UpdateImageStatus(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateImageStatus", reflect.TypeOf((*MockImageServiceInterface)(nil).UpdateImageStatus), image)
}

// SetErrorStatusOnImage mocks base method
func (m *MockImageServiceInterface) SetErrorStatusOnImage(err error, i *models.Image) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetErrorStatusOnImage", err, i)
}

// SetErrorStatusOnImage indicates an expected call of SetErrorStatusOnImage
func (mr *MockImageServiceInterfaceMockRecorder) SetErrorStatusOnImage(err, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetErrorStatusOnImage", reflect.TypeOf((*MockImageServiceInterface)(nil).SetErrorStatusOnImage), err, i)
}

// CreateRepoForImage mocks base method
func (m *MockImageServiceInterface) CreateRepoForImage(i *models.Image) (*models.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRepoForImage", i)
	ret0, _ := ret[0].(*models.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRepoForImage indicates an expected call of CreateRepoForImage
func (mr *MockImageServiceInterfaceMockRecorder) CreateRepoForImage(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRepoForImage", reflect.TypeOf((*MockImageServiceInterface)(nil).CreateRepoForImage), i)
}

// CreateInstallerForImage mocks base method
func (m *MockImageServiceInterface) CreateInstallerForImage(i *models.Image) (*models.Image, chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInstallerForImage", i)
	ret0, _ := ret[0].(*models.Image)
	ret1, _ := ret[1].(chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateInstallerForImage indicates an expected call of CreateInstallerForImage
func (mr *MockImageServiceInterfaceMockRecorder) CreateInstallerForImage(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInstallerForImage", reflect.TypeOf((*MockImageServiceInterface)(nil).CreateInstallerForImage), i)
}

// GetImageByID mocks base method
func (m *MockImageServiceInterface) GetImageByID(id string) (*models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageByID", id)
	ret0, _ := ret[0].(*models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageByID indicates an expected call of GetImageByID
func (mr *MockImageServiceInterfaceMockRecorder) GetImageByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageByID", reflect.TypeOf((*MockImageServiceInterface)(nil).GetImageByID), id)
}

// GetUpdateInfo mocks base method
func (m *MockImageServiceInterface) GetUpdateInfo(image models.Image) ([]models.ImageUpdateAvailable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdateInfo", image)
	ret0, _ := ret[0].([]models.ImageUpdateAvailable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpdateInfo indicates an expected call of GetUpdateInfo
func (mr *MockImageServiceInterfaceMockRecorder) GetUpdateInfo(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdateInfo", reflect.TypeOf((*MockImageServiceInterface)(nil).GetUpdateInfo), image)
}

// AddPackageInfo mocks base method
func (m *MockImageServiceInterface) AddPackageInfo(image *models.Image) (services.ImageDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPackageInfo", image)
	ret0, _ := ret[0].(services.ImageDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPackageInfo indicates an expected call of AddPackageInfo
func (mr *MockImageServiceInterfaceMockRecorder) AddPackageInfo(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPackageInfo", reflect.TypeOf((*MockImageServiceInterface)(nil).AddPackageInfo), image)
}

// GetImageByOSTreeCommitHash mocks base method
func (m *MockImageServiceInterface) GetImageByOSTreeCommitHash(commitHash string) (*models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageByOSTreeCommitHash", commitHash)
	ret0, _ := ret[0].(*models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageByOSTreeCommitHash indicates an expected call of GetImageByOSTreeCommitHash
func (mr *MockImageServiceInterfaceMockRecorder) GetImageByOSTreeCommitHash(commitHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageByOSTreeCommitHash", reflect.TypeOf((*MockImageServiceInterface)(nil).GetImageByOSTreeCommitHash), commitHash)
}

// CheckImageName mocks base method
func (m *MockImageServiceInterface) CheckImageName(name, account string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckImageName", name, account)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckImageName indicates an expected call of CheckImageName
func (mr *MockImageServiceInterfaceMockRecorder) CheckImageName(name, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckImageName", reflect.TypeOf((*MockImageServiceInterface)(nil).CheckImageName), name, account)
}

// RetryCreateImage mocks base method
func (m *MockImageServiceInterface) RetryCreateImage(image *models.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetryCreateImage", image)
	ret0, _ := ret[0].(error)
	return ret0
}

// RetryCreateImage indicates an expected call of RetryCreateImage
func (mr *MockImageServiceInterfaceMockRecorder) RetryCreateImage(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetryCreateImage", reflect.TypeOf((*MockImageServiceInterface)(nil).RetryCreateImage), image)
}

// GetMetadata mocks base method
func (m *MockImageServiceInterface) GetMetadata(image *models.Image) (*models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadata", image)
	ret0, _ := ret[0].(*models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadata indicates an expected call of GetMetadata
func (mr *MockImageServiceInterfaceMockRecorder) GetMetadata(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockImageServiceInterface)(nil).GetMetadata), image)
}

// SetFinalImageStatus mocks base method
func (m *MockImageServiceInterface) SetFinalImageStatus(i *models.Image) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFinalImageStatus", i)
}

// SetFinalImageStatus indicates an expected call of SetFinalImageStatus
func (mr *MockImageServiceInterfaceMockRecorder) SetFinalImageStatus(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFinalImageStatus", reflect.TypeOf((*MockImageServiceInterface)(nil).SetFinalImageStatus), i)
}

// CheckIfIsLatestVersion mocks base method
func (m *MockImageServiceInterface) CheckIfIsLatestVersion(previousImage *models.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIfIsLatestVersion", previousImage)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckIfIsLatestVersion indicates an expected call of CheckIfIsLatestVersion
func (mr *MockImageServiceInterfaceMockRecorder) CheckIfIsLatestVersion(previousImage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIfIsLatestVersion", reflect.TypeOf((*MockImageServiceInterface)(nil).CheckIfIsLatestVersion), previousImage)
}

// SetBuildingStatusOnImageToRetryBuild mocks base method
func (m *MockImageServiceInterface) SetBuildingStatusOnImageToRetryBuild(image *models.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBuildingStatusOnImageToRetryBuild", image)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBuildingStatusOnImageToRetryBuild indicates an expected call of SetBuildingStatusOnImageToRetryBuild
func (mr *MockImageServiceInterfaceMockRecorder) SetBuildingStatusOnImageToRetryBuild(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBuildingStatusOnImageToRetryBuild", reflect.TypeOf((*MockImageServiceInterface)(nil).SetBuildingStatusOnImageToRetryBuild), image)
}

// GetRollbackImage mocks base method
func (m *MockImageServiceInterface) GetRollbackImage(image *models.Image) (*models.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRollbackImage", image)
	ret0, _ := ret[0].(*models.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRollbackImage indicates an expected call of GetRollbackImage
func (mr *MockImageServiceInterfaceMockRecorder) GetRollbackImage(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRollbackImage", reflect.TypeOf((*MockImageServiceInterface)(nil).GetRollbackImage), image)
}

// SendImageNotification mocks base method
func (m *MockImageServiceInterface) SendImageNotification(image *models.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendImageNotification", image)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendImageNotification indicates an expected call of SendImageNotification
func (mr *MockImageServiceInterfaceMockRecorder) SendImageNotification(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendImageNotification", reflect.TypeOf((*MockImageServiceInterface)(nil).SendImageNotification), image)
}
