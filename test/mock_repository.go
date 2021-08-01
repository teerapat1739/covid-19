// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/teerapatthong-o/workspace/my/COVID-19/business/covid/repository.go

// Package mock_covid is a generated GoMock package.
package mock_covid

import (
	model "covid-19/business/covid/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetCovidCases mocks base method.
func (m *MockRepository) GetCovidCases() (model.ReponseCovidApiCases, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCovidCases")
	ret0, _ := ret[0].(model.ReponseCovidApiCases)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCovidCases indicates an expected call of GetCovidCases.
func (mr *MockRepositoryMockRecorder) GetCovidCases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCovidCases", reflect.TypeOf((*MockRepository)(nil).GetCovidCases))
}
