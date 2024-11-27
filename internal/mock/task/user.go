// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/task/domain/user.go
//
// Generated by this command:
//
//	mockgen -source=pkg/task/domain/user.go -destination=internal/mock/task/user.go -package=mocktask
//

// Package mocktask is a generated GoMock package.
package mocktask

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

// GetGoalQuota mocks base method.
func (m *MockUserHub) GetGoalQuota(ctx *appcontext.AppContext, userID string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGoalQuota", ctx, userID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGoalQuota indicates an expected call of GetGoalQuota.
func (mr *MockUserHubMockRecorder) GetGoalQuota(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGoalQuota", reflect.TypeOf((*MockUserHub)(nil).GetGoalQuota), ctx, userID)
}

// GetTaskQuota mocks base method.
func (m *MockUserHub) GetTaskQuota(ctx *appcontext.AppContext, userID string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskQuota", ctx, userID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskQuota indicates an expected call of GetTaskQuota.
func (mr *MockUserHubMockRecorder) GetTaskQuota(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskQuota", reflect.TypeOf((*MockUserHub)(nil).GetTaskQuota), ctx, userID)
}