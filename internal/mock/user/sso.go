// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/user/domain/sso.go
//
// Generated by this command:
//
//	mockgen -source=pkg/user/domain/sso.go -destination=internal/mock/user/sso.go -package=mockuser
//

// Package mockuser is a generated GoMock package.
package mockuser

import (
	reflect "reflect"

	appcontext "github.com/namhq1989/go-utilities/appcontext"
	domain "github.com/namhq1989/tapnchill-server/pkg/user/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockSSORepository is a mock of SSORepository interface.
type MockSSORepository struct {
	ctrl     *gomock.Controller
	recorder *MockSSORepositoryMockRecorder
}

// MockSSORepositoryMockRecorder is the mock recorder for MockSSORepository.
type MockSSORepositoryMockRecorder struct {
	mock *MockSSORepository
}

// NewMockSSORepository creates a new mock instance.
func NewMockSSORepository(ctrl *gomock.Controller) *MockSSORepository {
	mock := &MockSSORepository{ctrl: ctrl}
	mock.recorder = &MockSSORepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSSORepository) EXPECT() *MockSSORepositoryMockRecorder {
	return m.recorder
}

// VerifyGoogleToken mocks base method.
func (m *MockSSORepository) VerifyGoogleToken(ctx *appcontext.AppContext, token string) (*domain.SSOGoogleUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyGoogleToken", ctx, token)
	ret0, _ := ret[0].(*domain.SSOGoogleUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyGoogleToken indicates an expected call of VerifyGoogleToken.
func (mr *MockSSORepositoryMockRecorder) VerifyGoogleToken(ctx, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyGoogleToken", reflect.TypeOf((*MockSSORepository)(nil).VerifyGoogleToken), ctx, token)
}
