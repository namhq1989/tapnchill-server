// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/habit/domain/habit.go
//
// Generated by this command:
//
//	mockgen -source=pkg/habit/domain/habit.go -destination=internal/mock/habit/habit.go -package=mockhabit
//

// Package mockhabit is a generated GoMock package.
package mockhabit

import (
	reflect "reflect"
	time "time"

	appcontext "github.com/namhq1989/go-utilities/appcontext"
	domain "github.com/namhq1989/tapnchill-server/pkg/habit/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockHabitRepository is a mock of HabitRepository interface.
type MockHabitRepository struct {
	ctrl     *gomock.Controller
	recorder *MockHabitRepositoryMockRecorder
}

// MockHabitRepositoryMockRecorder is the mock recorder for MockHabitRepository.
type MockHabitRepositoryMockRecorder struct {
	mock *MockHabitRepository
}

// NewMockHabitRepository creates a new mock instance.
func NewMockHabitRepository(ctrl *gomock.Controller) *MockHabitRepository {
	mock := &MockHabitRepository{ctrl: ctrl}
	mock.recorder = &MockHabitRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHabitRepository) EXPECT() *MockHabitRepositoryMockRecorder {
	return m.recorder
}

// CountScheduledHabits mocks base method.
func (m *MockHabitRepository) CountScheduledHabits(ctx *appcontext.AppContext, userID string, date time.Time) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountScheduledHabits", ctx, userID, date)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountScheduledHabits indicates an expected call of CountScheduledHabits.
func (mr *MockHabitRepositoryMockRecorder) CountScheduledHabits(ctx, userID, date any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountScheduledHabits", reflect.TypeOf((*MockHabitRepository)(nil).CountScheduledHabits), ctx, userID, date)
}

// Create mocks base method.
func (m *MockHabitRepository) Create(ctx *appcontext.AppContext, habit domain.Habit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, habit)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockHabitRepositoryMockRecorder) Create(ctx, habit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockHabitRepository)(nil).Create), ctx, habit)
}

// Delete mocks base method.
func (m *MockHabitRepository) Delete(ctx *appcontext.AppContext, habitID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, habitID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockHabitRepositoryMockRecorder) Delete(ctx, habitID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockHabitRepository)(nil).Delete), ctx, habitID)
}

// FindByFilter mocks base method.
func (m *MockHabitRepository) FindByFilter(ctx *appcontext.AppContext, filter domain.HabitFilter) ([]domain.Habit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByFilter", ctx, filter)
	ret0, _ := ret[0].([]domain.Habit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByFilter indicates an expected call of FindByFilter.
func (mr *MockHabitRepositoryMockRecorder) FindByFilter(ctx, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByFilter", reflect.TypeOf((*MockHabitRepository)(nil).FindByFilter), ctx, filter)
}

// FindByID mocks base method.
func (m *MockHabitRepository) FindByID(ctx *appcontext.AppContext, habitID string) (*domain.Habit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, habitID)
	ret0, _ := ret[0].(*domain.Habit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockHabitRepositoryMockRecorder) FindByID(ctx, habitID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockHabitRepository)(nil).FindByID), ctx, habitID)
}

// Update mocks base method.
func (m *MockHabitRepository) Update(ctx *appcontext.AppContext, habit domain.Habit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, habit)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockHabitRepositoryMockRecorder) Update(ctx, habit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockHabitRepository)(nil).Update), ctx, habit)
}