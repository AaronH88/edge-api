// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/devices.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	gomock "github.com/golang/mock/gomock"
	inventory "github.com/redhatinsights/edge-api/pkg/clients/inventory"
	models "github.com/redhatinsights/edge-api/pkg/models"
	reflect "reflect"
)

// MockDeviceServiceInterface is a mock of DeviceServiceInterface interface
type MockDeviceServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceServiceInterfaceMockRecorder
}

// MockDeviceServiceInterfaceMockRecorder is the mock recorder for MockDeviceServiceInterface
type MockDeviceServiceInterfaceMockRecorder struct {
	mock *MockDeviceServiceInterface
}

// NewMockDeviceServiceInterface creates a new mock instance
func NewMockDeviceServiceInterface(ctrl *gomock.Controller) *MockDeviceServiceInterface {
	mock := &MockDeviceServiceInterface{ctrl: ctrl}
	mock.recorder = &MockDeviceServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeviceServiceInterface) EXPECT() *MockDeviceServiceInterfaceMockRecorder {
	return m.recorder
}

// GetDevices mocks base method
func (m *MockDeviceServiceInterface) GetDevices(params *inventory.Params) (*models.DeviceDetailsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevices", params)
	ret0, _ := ret[0].(*models.DeviceDetailsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevices indicates an expected call of GetDevices
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDevices(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevices", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDevices), params)
}

// GetDevicesView mocks base method
func (m *MockDeviceServiceInterface) GetDevicesView(params *inventory.Params) (*models.DeviceViewList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevicesView", params)
	ret0, _ := ret[0].(*models.DeviceViewList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevicesView indicates an expected call of GetDevicesView
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDevicesView(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevicesView", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDevicesView), params)
}

// GetDeviceByID mocks base method
func (m *MockDeviceServiceInterface) GetDeviceByID(deviceID uint) (*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByID", deviceID)
	ret0, _ := ret[0].(*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceByID indicates an expected call of GetDeviceByID
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceByID(deviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByID", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceByID), deviceID)
}

// GetDeviceByUUID mocks base method
func (m *MockDeviceServiceInterface) GetDeviceByUUID(deviceUUID string) (*models.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceByUUID", deviceUUID)
	ret0, _ := ret[0].(*models.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceByUUID indicates an expected call of GetDeviceByUUID
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceByUUID(deviceUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceByUUID", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceByUUID), deviceUUID)
}

// GetDeviceDetailsByUUID mocks base method
func (m *MockDeviceServiceInterface) GetDeviceDetailsByUUID(deviceUUID string) (*models.DeviceDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceDetailsByUUID", deviceUUID)
	ret0, _ := ret[0].(*models.DeviceDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceDetailsByUUID indicates an expected call of GetDeviceDetailsByUUID
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceDetailsByUUID(deviceUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceDetailsByUUID", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceDetailsByUUID), deviceUUID)
}

// GetUpdateAvailableForDeviceByUUID mocks base method
func (m *MockDeviceServiceInterface) GetUpdateAvailableForDeviceByUUID(deviceUUID string) ([]models.ImageUpdateAvailable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdateAvailableForDeviceByUUID", deviceUUID)
	ret0, _ := ret[0].([]models.ImageUpdateAvailable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpdateAvailableForDeviceByUUID indicates an expected call of GetUpdateAvailableForDeviceByUUID
func (mr *MockDeviceServiceInterfaceMockRecorder) GetUpdateAvailableForDeviceByUUID(deviceUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdateAvailableForDeviceByUUID", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetUpdateAvailableForDeviceByUUID), deviceUUID)
}

// GetDeviceImageInfoByUUID mocks base method
func (m *MockDeviceServiceInterface) GetDeviceImageInfoByUUID(deviceUUID string) (*models.ImageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceImageInfoByUUID", deviceUUID)
	ret0, _ := ret[0].(*models.ImageInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceImageInfoByUUID indicates an expected call of GetDeviceImageInfoByUUID
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceImageInfoByUUID(deviceUUID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceImageInfoByUUID", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceImageInfoByUUID), deviceUUID)
}

// GetDeviceDetails mocks base method
func (m *MockDeviceServiceInterface) GetDeviceDetails(device inventory.Device) (*models.DeviceDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceDetails", device)
	ret0, _ := ret[0].(*models.DeviceDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceDetails indicates an expected call of GetDeviceDetails
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceDetails(device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceDetails", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceDetails), device)
}

// GetUpdateAvailableForDevice mocks base method
func (m *MockDeviceServiceInterface) GetUpdateAvailableForDevice(device inventory.Device) ([]models.ImageUpdateAvailable, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdateAvailableForDevice", device)
	ret0, _ := ret[0].([]models.ImageUpdateAvailable)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpdateAvailableForDevice indicates an expected call of GetUpdateAvailableForDevice
func (mr *MockDeviceServiceInterfaceMockRecorder) GetUpdateAvailableForDevice(device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdateAvailableForDevice", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetUpdateAvailableForDevice), device)
}

// GetDeviceImageInfo mocks base method
func (m *MockDeviceServiceInterface) GetDeviceImageInfo(device inventory.Device) (*models.ImageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceImageInfo", device)
	ret0, _ := ret[0].(*models.ImageInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeviceImageInfo indicates an expected call of GetDeviceImageInfo
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceImageInfo(device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceImageInfo", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceImageInfo), device)
}

// GetDeviceLastDeployment mocks base method
func (m *MockDeviceServiceInterface) GetDeviceLastDeployment(device inventory.Device) *inventory.OSTree {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceLastDeployment", device)
	ret0, _ := ret[0].(*inventory.OSTree)
	return ret0
}

// GetDeviceLastDeployment indicates an expected call of GetDeviceLastDeployment
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceLastDeployment(device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceLastDeployment", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceLastDeployment), device)
}

// GetDeviceLastBootedDeployment mocks base method
func (m *MockDeviceServiceInterface) GetDeviceLastBootedDeployment(device inventory.Device) *inventory.OSTree {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeviceLastBootedDeployment", device)
	ret0, _ := ret[0].(*inventory.OSTree)
	return ret0
}

// GetDeviceLastBootedDeployment indicates an expected call of GetDeviceLastBootedDeployment
func (mr *MockDeviceServiceInterfaceMockRecorder) GetDeviceLastBootedDeployment(device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeviceLastBootedDeployment", reflect.TypeOf((*MockDeviceServiceInterface)(nil).GetDeviceLastBootedDeployment), device)
}

// ProcessPlatformInventoryCreateEvent mocks base method
func (m *MockDeviceServiceInterface) ProcessPlatformInventoryCreateEvent(message []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessPlatformInventoryCreateEvent", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessPlatformInventoryCreateEvent indicates an expected call of ProcessPlatformInventoryCreateEvent
func (mr *MockDeviceServiceInterfaceMockRecorder) ProcessPlatformInventoryCreateEvent(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessPlatformInventoryCreateEvent", reflect.TypeOf((*MockDeviceServiceInterface)(nil).ProcessPlatformInventoryCreateEvent), message)
}

// ProcessPlatformInventoryUpdatedEvent mocks base method
func (m *MockDeviceServiceInterface) ProcessPlatformInventoryUpdatedEvent(message []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessPlatformInventoryUpdatedEvent", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessPlatformInventoryUpdatedEvent indicates an expected call of ProcessPlatformInventoryUpdatedEvent
func (mr *MockDeviceServiceInterfaceMockRecorder) ProcessPlatformInventoryUpdatedEvent(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessPlatformInventoryUpdatedEvent", reflect.TypeOf((*MockDeviceServiceInterface)(nil).ProcessPlatformInventoryUpdatedEvent), message)
}

// ProcessPlatformInventoryDeleteEvent mocks base method
func (m *MockDeviceServiceInterface) ProcessPlatformInventoryDeleteEvent(message []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessPlatformInventoryDeleteEvent", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessPlatformInventoryDeleteEvent indicates an expected call of ProcessPlatformInventoryDeleteEvent
func (mr *MockDeviceServiceInterfaceMockRecorder) ProcessPlatformInventoryDeleteEvent(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessPlatformInventoryDeleteEvent", reflect.TypeOf((*MockDeviceServiceInterface)(nil).ProcessPlatformInventoryDeleteEvent), message)
}
