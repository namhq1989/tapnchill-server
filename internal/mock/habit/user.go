// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/habit/domain/user.go
//
// Generated by this command:
//
//	mockgen -source=pkg/habit/domain/user.go -destination=internal/mock/habit/user.go -package=mockhabit
//

// Package mockhabit is a generated GoMock package.
package mockhabit

import (
	reflect "reflect"

	appcontext "github.com/namhq1989/go-utilities/appcontext"
	gomock "go.uber.org/mock/gomock"
)

// MockUserHub is a mock of UserHub interface.
type MockUserHub struct {
	ctrl     *gomock.Controller
	recorder *MockUserHubMockRecorder
}

// MockUserHubMockRecorder is the mock recorder for MockUserHub.
type MockUserHubMockRecorder struct {
	mock *MockUserHub
}

// NewMockUserHub creates a new mock instance.
func NewMockUserHub(ctrl *gomock.Controller) *MockUserHub {
	mock := &MockUserHub{ctrl: ctrl}
	mock.recorder = &MockUserHubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserHub) EXPECT() *MockUserHubMockRecorder {
	return m.recorder
}

// GetHabitQuota mocks base method.
func (m *MockUserHub) GetHabitQuota(ctx *appcontext.AppContext, userID string) (int64, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHabitQuota", ctx, userID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetHabitQuota indicates an expected call of GetHabitQuota.
func (mr *MockUserHubMockRecorder) GetHabitQuota(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHabitQuota", reflect.TypeOf((*MockUserHub)(nil).GetHabitQuota), ctx, userID)
}
