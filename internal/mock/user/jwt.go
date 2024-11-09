// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/user/domain/jwt.go
//
// Generated by this command:
//
//	mockgen -source=pkg/user/domain/jwt.go -destination=internal/mock/user/jwt.go -package=mockuser
//

// Package mockuser is a generated GoMock package.
package mockuser

import (
	reflect "reflect"

	appcontext "github.com/namhq1989/go-utilities/appcontext"
	gomock "go.uber.org/mock/gomock"
)

// MockJwtRepository is a mock of JwtRepository interface.
type MockJwtRepository struct {
	ctrl     *gomock.Controller
	recorder *MockJwtRepositoryMockRecorder
}

// MockJwtRepositoryMockRecorder is the mock recorder for MockJwtRepository.
type MockJwtRepositoryMockRecorder struct {
	mock *MockJwtRepository
}

// NewMockJwtRepository creates a new mock instance.
func NewMockJwtRepository(ctrl *gomock.Controller) *MockJwtRepository {
	mock := &MockJwtRepository{ctrl: ctrl}
	mock.recorder = &MockJwtRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJwtRepository) EXPECT() *MockJwtRepositoryMockRecorder {
	return m.recorder
}

// GenerateAccessToken mocks base method.
func (m *MockJwtRepository) GenerateAccessToken(ctx *appcontext.AppContext, userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", ctx, userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockJwtRepositoryMockRecorder) GenerateAccessToken(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockJwtRepository)(nil).GenerateAccessToken), ctx, userID)
}