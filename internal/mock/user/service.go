// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/user/domain/service.go
//
// Generated by this command:
//
//	mockgen -source=pkg/user/domain/service.go -destination=internal/mock/user/service.go -package=mockuser
//

// Package mockuser is a generated GoMock package.
package mockuser

import (
	reflect "reflect"

	appcontext "github.com/namhq1989/go-utilities/appcontext"
	domain "github.com/namhq1989/tapnchill-server/pkg/user/domain"
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

// GetLemonsqueezyCustomerPortalURL mocks base method.
func (m *MockService) GetLemonsqueezyCustomerPortalURL(ctx *appcontext.AppContext, userID string) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLemonsqueezyCustomerPortalURL", ctx, userID)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLemonsqueezyCustomerPortalURL indicates an expected call of GetLemonsqueezyCustomerPortalURL.
func (mr *MockServiceMockRecorder) GetLemonsqueezyCustomerPortalURL(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLemonsqueezyCustomerPortalURL", reflect.TypeOf((*MockService)(nil).GetLemonsqueezyCustomerPortalURL), ctx, userID)
}

// GetUserByID mocks base method.
func (m *MockService) GetUserByID(ctx *appcontext.AppContext, userID string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, userID)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockServiceMockRecorder) GetUserByID(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockService)(nil).GetUserByID), ctx, userID)
}
