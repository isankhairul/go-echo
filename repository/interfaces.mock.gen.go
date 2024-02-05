// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=repository/interfaces.go -destination=repository/interfaces.mock.gen.go -package=repository
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	entity "go-echo/model/entity"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetTestById mocks base method.
func (m *MockRepositoryInterface) GetTestById(ctx context.Context, input GetTestByIdInput) (GetTestByIdOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestById", ctx, input)
	ret0, _ := ret[0].(GetTestByIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestById indicates an expected call of GetTestById.
func (mr *MockRepositoryInterfaceMockRecorder) GetTestById(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetTestById), ctx, input)
}

// UsersCountByPhone mocks base method.
func (m *MockRepositoryInterface) UsersCountByPhone(ctx context.Context, phone string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsersCountByPhone", ctx, phone)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsersCountByPhone indicates an expected call of UsersCountByPhone.
func (mr *MockRepositoryInterfaceMockRecorder) UsersCountByPhone(ctx, phone any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsersCountByPhone", reflect.TypeOf((*MockRepositoryInterface)(nil).UsersCountByPhone), ctx, phone)
}

// UsersCreate mocks base method.
func (m *MockRepositoryInterface) UsersCreate(ctx context.Context, Users entity.Users) (*int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsersCreate", ctx, Users)
	ret0, _ := ret[0].(*int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsersCreate indicates an expected call of UsersCreate.
func (mr *MockRepositoryInterfaceMockRecorder) UsersCreate(ctx, Users any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsersCreate", reflect.TypeOf((*MockRepositoryInterface)(nil).UsersCreate), ctx, Users)
}

// UsersFirstByID mocks base method.
func (m *MockRepositoryInterface) UsersFirstByID(ctx context.Context, id int64) (*entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsersFirstByID", ctx, id)
	ret0, _ := ret[0].(*entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsersFirstByID indicates an expected call of UsersFirstByID.
func (mr *MockRepositoryInterfaceMockRecorder) UsersFirstByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsersFirstByID", reflect.TypeOf((*MockRepositoryInterface)(nil).UsersFirstByID), ctx, id)
}

// UsersFirstByPhone mocks base method.
func (m *MockRepositoryInterface) UsersFirstByPhone(ctx context.Context, phone string) (*entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsersFirstByPhone", ctx, phone)
	ret0, _ := ret[0].(*entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsersFirstByPhone indicates an expected call of UsersFirstByPhone.
func (mr *MockRepositoryInterfaceMockRecorder) UsersFirstByPhone(ctx, phone any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsersFirstByPhone", reflect.TypeOf((*MockRepositoryInterface)(nil).UsersFirstByPhone), ctx, phone)
}

// UsersUpdateByID mocks base method.
func (m *MockRepositoryInterface) UsersUpdateByID(ctx context.Context, id int64, Users entity.Users) (*int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsersUpdateByID", ctx, id, Users)
	ret0, _ := ret[0].(*int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UsersUpdateByID indicates an expected call of UsersUpdateByID.
func (mr *MockRepositoryInterfaceMockRecorder) UsersUpdateByID(ctx, id, Users any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsersUpdateByID", reflect.TypeOf((*MockRepositoryInterface)(nil).UsersUpdateByID), ctx, id, Users)
}
