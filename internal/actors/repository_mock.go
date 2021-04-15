// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_actors is a generated GoMock package.
package actors

import (
	reflect "reflect"

	models "github.com/go-park-mail-ru/2021_1_kekEnd/internal/models"
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

// CreateActor mocks base method.
func (m *MockRepository) CreateActor(arg0 models.Actor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateActor", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateActor indicates an expected call of CreateActor.
func (mr *MockRepositoryMockRecorder) CreateActor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateActor", reflect.TypeOf((*MockRepository)(nil).CreateActor), arg0)
}

// EditActor mocks base method.
func (m *MockRepository) EditActor(arg0 models.Actor) (models.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditActor", arg0)
	ret0, _ := ret[0].(models.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditActor indicates an expected call of EditActor.
func (mr *MockRepositoryMockRecorder) EditActor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditActor", reflect.TypeOf((*MockRepository)(nil).EditActor), arg0)
}

// GetActorByID mocks base method.
func (m *MockRepository) GetActorByID(id string) (models.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorByID", id)
	ret0, _ := ret[0].(models.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorByID indicates an expected call of GetActorByID.
func (mr *MockRepositoryMockRecorder) GetActorByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorByID", reflect.TypeOf((*MockRepository)(nil).GetActorByID), id)
}
