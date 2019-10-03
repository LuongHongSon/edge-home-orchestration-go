// Code generated by MockGen. DO NOT EDIT.
// Source: orchestration.go

// Package mock_orchestrationapi is a generated GoMock package.
package mocks

import (
	configuremgrtypes "common/types/configuremgrtypes"
	gomock "github.com/golang/mock/gomock"
	orchestrationapi "orchestrationapi"
	reflect "reflect"
)

// MockOrche is a mock of Orche interface
type MockOrche struct {
	ctrl     *gomock.Controller
	recorder *MockOrcheMockRecorder
}

// MockOrcheMockRecorder is the mock recorder for MockOrche
type MockOrcheMockRecorder struct {
	mock *MockOrche
}

// NewMockOrche creates a new mock instance
func NewMockOrche(ctrl *gomock.Controller) *MockOrche {
	mock := &MockOrche{ctrl: ctrl}
	mock.recorder = &MockOrcheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrche) EXPECT() *MockOrcheMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockOrche) Start(deviceIDPath, platform, executionType string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", deviceIDPath, platform, executionType)
}

// Start indicates an expected call of Start
func (mr *MockOrcheMockRecorder) Start(deviceIDPath, platform, executionType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockOrche)(nil).Start), deviceIDPath, platform, executionType)
}

// MockOrcheExternalAPI is a mock of OrcheExternalAPI interface
type MockOrcheExternalAPI struct {
	ctrl     *gomock.Controller
	recorder *MockOrcheExternalAPIMockRecorder
}

// MockOrcheExternalAPIMockRecorder is the mock recorder for MockOrcheExternalAPI
type MockOrcheExternalAPIMockRecorder struct {
	mock *MockOrcheExternalAPI
}

// NewMockOrcheExternalAPI creates a new mock instance
func NewMockOrcheExternalAPI(ctrl *gomock.Controller) *MockOrcheExternalAPI {
	mock := &MockOrcheExternalAPI{ctrl: ctrl}
	mock.recorder = &MockOrcheExternalAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrcheExternalAPI) EXPECT() *MockOrcheExternalAPIMockRecorder {
	return m.recorder
}

// RequestService mocks base method
func (m *MockOrcheExternalAPI) RequestService(serviceInfo orchestrationapi.ReqeustService) orchestrationapi.ResponseService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestService", serviceInfo)
	ret0, _ := ret[0].(orchestrationapi.ResponseService)
	return ret0
}

// RequestService indicates an expected call of RequestService
func (mr *MockOrcheExternalAPIMockRecorder) RequestService(serviceInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestService", reflect.TypeOf((*MockOrcheExternalAPI)(nil).RequestService), serviceInfo)
}

// MockOrcheInternalAPI is a mock of OrcheInternalAPI interface
type MockOrcheInternalAPI struct {
	ctrl     *gomock.Controller
	recorder *MockOrcheInternalAPIMockRecorder
}

// MockOrcheInternalAPIMockRecorder is the mock recorder for MockOrcheInternalAPI
type MockOrcheInternalAPIMockRecorder struct {
	mock *MockOrcheInternalAPI
}

// NewMockOrcheInternalAPI creates a new mock instance
func NewMockOrcheInternalAPI(ctrl *gomock.Controller) *MockOrcheInternalAPI {
	mock := &MockOrcheInternalAPI{ctrl: ctrl}
	mock.recorder = &MockOrcheInternalAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrcheInternalAPI) EXPECT() *MockOrcheInternalAPIMockRecorder {
	return m.recorder
}

// Notify mocks base method
func (m *MockOrcheInternalAPI) Notify(serviceinfo configuremgrtypes.ServiceInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Notify", serviceinfo)
}

// Notify indicates an expected call of Notify
func (mr *MockOrcheInternalAPIMockRecorder) Notify(serviceinfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notify", reflect.TypeOf((*MockOrcheInternalAPI)(nil).Notify), serviceinfo)
}

// ExecuteAppOnLocal mocks base method
func (m *MockOrcheInternalAPI) ExecuteAppOnLocal(appInfo map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExecuteAppOnLocal", appInfo)
}

// ExecuteAppOnLocal indicates an expected call of ExecuteAppOnLocal
func (mr *MockOrcheInternalAPIMockRecorder) ExecuteAppOnLocal(appInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteAppOnLocal", reflect.TypeOf((*MockOrcheInternalAPI)(nil).ExecuteAppOnLocal), appInfo)
}

// HandleNotificationOnLocal mocks base method
func (m *MockOrcheInternalAPI) HandleNotificationOnLocal(serviceID float64, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleNotificationOnLocal", serviceID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleNotificationOnLocal indicates an expected call of HandleNotificationOnLocal
func (mr *MockOrcheInternalAPIMockRecorder) HandleNotificationOnLocal(serviceID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleNotificationOnLocal", reflect.TypeOf((*MockOrcheInternalAPI)(nil).HandleNotificationOnLocal), serviceID, status)
}

// GetScore mocks base method
func (m *MockOrcheInternalAPI) GetScore(target string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScore", target)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScore indicates an expected call of GetScore
func (mr *MockOrcheInternalAPIMockRecorder) GetScore(target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScore", reflect.TypeOf((*MockOrcheInternalAPI)(nil).GetScore), target)
}
