// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/repository/user_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"
	graph "github.com/chaki8923/wedding-backend/pkg/domain/model/graph"
	gomock "github.com/golang/mock/gomock"
	model "github.com/chaki8923/wedding-backend/pkg/domain/model"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// CreateMessage mocks base method.
func (m *MockUser) CreateUser(user *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMessage indicates an expected call of CreateMessage.
func (mr *MockUserMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUser)(nil).CreateUser), user)
}

// Decrypt mocks base method.
func (m *MockUser) Decrypt(encrypted string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decrypt", encrypted)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decrypt indicates an expected call of Decrypt.
func (mr *MockUserMockRecorder) Decrypt(encrypted interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrypt", reflect.TypeOf((*MockUser)(nil).Decrypt), encrypted)
}

// Encrypt mocks base method.
func (m *MockUser) Encrypt(plain string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", plain)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt.
func (mr *MockUserMockRecorder) Encrypt(plain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockUser)(nil).Encrypt), plain)
}

// GetMapInIDs mocks base method.
func (m *MockUser) GetMapInIDs(ctx context.Context, ids []string) (map[string]*graph.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMapInIDs", ctx, ids)
	ret0, _ := ret[0].(map[string]*graph.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMapInIDs indicates an expected call of GetMapInIDs.
func (mr *MockUserMockRecorder) GetMapInIDs(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMapInIDs", reflect.TypeOf((*MockUser)(nil).GetMapInIDs), ctx, ids)
}

// GetUserByEmail mocks base method.
func (m *MockUser) GetUserByEmail(email string) (*graph.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(*graph.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUser)(nil).GetUserByEmail), email)
}

// GetUserById mocks base method.
func (m *MockUser) GetUserById(id string) (*graph.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", id)
	ret0, _ := ret[0].(*graph.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockUserMockRecorder) GetUserById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockUser)(nil).GetUserById), id)
}
