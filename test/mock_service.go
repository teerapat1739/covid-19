// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/teerapatthong-o/workspace/my/COVID-19/business/covid/service.go

// Package mock_covid is a generated GoMock package.
package mock_covid

import (
	model "covid-19/business/covid/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CovidSummary mocks base method.
func (m *MockService) CovidSummary() (model.Summary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CovidSummary")
	ret0, _ := ret[0].(model.Summary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CovidSummary indicates an expected call of CovidSummary.
func (mr *MockServiceMockRecorder) CovidSummary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CovidSummary", reflect.TypeOf((*MockService)(nil).CovidSummary))
}
