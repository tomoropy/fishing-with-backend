// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tomoropy/fishing-with-backend/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/tomoropy/fishing-with-backend/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateInvitation mocks base method.
func (m *MockStore) CreateInvitation(arg0 context.Context, arg1 db.CreateInvitationParams) (db.Invitations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInvitation", arg0, arg1)
	ret0, _ := ret[0].(db.Invitations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInvitation indicates an expected call of CreateInvitation.
func (mr *MockStoreMockRecorder) CreateInvitation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvitation", reflect.TypeOf((*MockStore)(nil).CreateInvitation), arg0, arg1)
}

// CreatePhoto mocks base method.
func (m *MockStore) CreatePhoto(arg0 context.Context, arg1 db.CreatePhotoParams) (db.Photos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePhoto", arg0, arg1)
	ret0, _ := ret[0].(db.Photos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePhoto indicates an expected call of CreatePhoto.
func (mr *MockStoreMockRecorder) CreatePhoto(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePhoto", reflect.TypeOf((*MockStore)(nil).CreatePhoto), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteInvitation mocks base method.
func (m *MockStore) DeleteInvitation(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInvitation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInvitation indicates an expected call of DeleteInvitation.
func (mr *MockStoreMockRecorder) DeleteInvitation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInvitation", reflect.TypeOf((*MockStore)(nil).DeleteInvitation), arg0, arg1)
}

// DeletePhoto mocks base method.
func (m *MockStore) DeletePhoto(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePhoto", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePhoto indicates an expected call of DeletePhoto.
func (mr *MockStoreMockRecorder) DeletePhoto(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhoto", reflect.TypeOf((*MockStore)(nil).DeletePhoto), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// GetInvitation mocks base method.
func (m *MockStore) GetInvitation(arg0 context.Context, arg1 int32) (db.Invitations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInvitation", arg0, arg1)
	ret0, _ := ret[0].(db.Invitations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInvitation indicates an expected call of GetInvitation.
func (mr *MockStoreMockRecorder) GetInvitation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvitation", reflect.TypeOf((*MockStore)(nil).GetInvitation), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 int32) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserByName mocks base method.
func (m *MockStore) GetUserByName(arg0 context.Context, arg1 string) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockStoreMockRecorder) GetUserByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockStore)(nil).GetUserByName), arg0, arg1)
}

// ListInvitation mocks base method.
func (m *MockStore) ListInvitation(arg0 context.Context, arg1 db.ListInvitationParams) ([]db.Invitations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInvitation", arg0, arg1)
	ret0, _ := ret[0].([]db.Invitations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInvitation indicates an expected call of ListInvitation.
func (mr *MockStoreMockRecorder) ListInvitation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInvitation", reflect.TypeOf((*MockStore)(nil).ListInvitation), arg0, arg1)
}

// ListInvitationByUser mocks base method.
func (m *MockStore) ListInvitationByUser(arg0 context.Context, arg1 db.ListInvitationByUserParams) ([]db.Invitations, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInvitationByUser", arg0, arg1)
	ret0, _ := ret[0].([]db.Invitations)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInvitationByUser indicates an expected call of ListInvitationByUser.
func (mr *MockStoreMockRecorder) ListInvitationByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInvitationByUser", reflect.TypeOf((*MockStore)(nil).ListInvitationByUser), arg0, arg1)
}

// ListUser mocks base method.
func (m *MockStore) ListUser(arg0 context.Context, arg1 db.ListUserParams) ([]db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUser", arg0, arg1)
	ret0, _ := ret[0].([]db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUser indicates an expected call of ListUser.
func (mr *MockStoreMockRecorder) ListUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUser", reflect.TypeOf((*MockStore)(nil).ListUser), arg0, arg1)
}

// UpdateInvitation mocks base method.
func (m *MockStore) UpdateInvitation(arg0 context.Context, arg1 db.UpdateInvitationParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInvitation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInvitation indicates an expected call of UpdateInvitation.
func (mr *MockStoreMockRecorder) UpdateInvitation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInvitation", reflect.TypeOf((*MockStore)(nil).UpdateInvitation), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}

// UserPhotoList mocks base method.
func (m *MockStore) UserPhotoList(arg0 context.Context, arg1 db.UserPhotoListParams) ([]db.Photos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserPhotoList", arg0, arg1)
	ret0, _ := ret[0].([]db.Photos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserPhotoList indicates an expected call of UserPhotoList.
func (mr *MockStoreMockRecorder) UserPhotoList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserPhotoList", reflect.TypeOf((*MockStore)(nil).UserPhotoList), arg0, arg1)
}
