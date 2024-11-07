// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/task/domain/goal.go
//
// Generated by this command:
//
//	mockgen -source=pkg/task/domain/goal.go -destination=internal/mock/task/goal.go -package=mocktask
//

// Package mocktask is a generated GoMock package.
package mocktask

import (
	reflect "reflect"

	appcontext "github.com/namhq1989/go-utilities/appcontext"
	domain "github.com/namhq1989/tapnchill-server/pkg/task/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockGoalRepository is a mock of GoalRepository interface.
type MockGoalRepository struct {
	ctrl     *gomock.Controller
	recorder *MockGoalRepositoryMockRecorder
}

// MockGoalRepositoryMockRecorder is the mock recorder for MockGoalRepository.
type MockGoalRepositoryMockRecorder struct {
	mock *MockGoalRepository
}

// NewMockGoalRepository creates a new mock instance.
func NewMockGoalRepository(ctrl *gomock.Controller) *MockGoalRepository {
	mock := &MockGoalRepository{ctrl: ctrl}
	mock.recorder = &MockGoalRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGoalRepository) EXPECT() *MockGoalRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockGoalRepository) Create(ctx *appcontext.AppContext, goal domain.Goal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, goal)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockGoalRepositoryMockRecorder) Create(ctx, goal any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGoalRepository)(nil).Create), ctx, goal)
}

// FindByFilter mocks base method.
func (m *MockGoalRepository) FindByFilter(ctx *appcontext.AppContext, filter domain.GoalFilter) ([]domain.Goal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByFilter", ctx, filter)
	ret0, _ := ret[0].([]domain.Goal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByFilter indicates an expected call of FindByFilter.
func (mr *MockGoalRepositoryMockRecorder) FindByFilter(ctx, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByFilter", reflect.TypeOf((*MockGoalRepository)(nil).FindByFilter), ctx, filter)
}

// Update mocks base method.
func (m *MockGoalRepository) Update(ctx *appcontext.AppContext, goal domain.Goal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, goal)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockGoalRepositoryMockRecorder) Update(ctx, goal any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGoalRepository)(nil).Update), ctx, goal)
}