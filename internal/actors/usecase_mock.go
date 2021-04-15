// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_actors is a generated GoMock package.
package actors

import (
	reflect "reflect"

	models "github.com/go-park-mail-ru/2021_1_kekEnd/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// CreateActor mocks base method.
func (m *MockUseCase) CreateActor(user models.User, actor models.Actor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateActor", user, actor)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateActor indicates an expected call of CreateActor.
func (mr *MockUseCaseMockRecorder) CreateActor(user, actor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateActor", reflect.TypeOf((*MockUseCase)(nil).CreateActor), user, actor)
}

// EditActor mocks base method.
func (m *MockUseCase) EditActor(user models.User, change models.Actor) (models.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditActor", user, change)
	ret0, _ := ret[0].(models.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditActor indicates an expected call of EditActor.
func (mr *MockUseCaseMockRecorder) EditActor(user, change interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditActor", reflect.TypeOf((*MockUseCase)(nil).EditActor), user, change)
}

// GetActor mocks base method.
func (m *MockUseCase) GetActor(id string) (models.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActor", id)
	ret0, _ := ret[0].(models.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActor indicates an expected call of GetActor.
func (mr *MockUseCaseMockRecorder) GetActor(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActor", reflect.TypeOf((*MockUseCase)(nil).GetActor), id)
}
