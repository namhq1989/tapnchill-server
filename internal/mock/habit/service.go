// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/habit/domain/service.go
//
// Generated by this command:
//
//	mockgen -source=pkg/habit/domain/service.go -destination=internal/mock/habit/service.go -package=mockhabit
//

// Package mockhabit is a generated GoMock package.
package mockhabit

import (
	reflect "reflect"

	appcontext "github.com/namhq1989/go-utilities/appcontext"
	domain "github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	gomock "go.uber.org/mock/gomock"
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

// GetHabitByID mocks base method.
func (m *MockService) GetHabitByID(ctx *appcontext.AppContext, habitID, userID string) (*domain.Habit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHabitByID", ctx, habitID, userID)
	ret0, _ := ret[0].(*domain.Habit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHabitByID indicates an expected call of GetHabitByID.
func (mr *MockServiceMockRecorder) GetHabitByID(ctx, habitID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHabitByID", reflect.TypeOf((*MockService)(nil).GetHabitByID), ctx, habitID, userID)
}