// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/common/domain/quote.go
//
// Generated by this command:
//
//	mockgen -source=pkg/common/domain/quote.go -destination=internal/mock/common/quote.go -package=mockcommon
//

// Package mockcommon is a generated GoMock package.
package mockcommon

import (
	reflect "reflect"

	appcontext "github.com/namhq1989/go-utilities/appcontext"
	domain "github.com/namhq1989/tapnchill-server/pkg/common/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockQuoteRepository is a mock of QuoteRepository interface.
type MockQuoteRepository struct {
	ctrl     *gomock.Controller
	recorder *MockQuoteRepositoryMockRecorder
}

// MockQuoteRepositoryMockRecorder is the mock recorder for MockQuoteRepository.
type MockQuoteRepositoryMockRecorder struct {
	mock *MockQuoteRepository
}

// NewMockQuoteRepository creates a new mock instance.
func NewMockQuoteRepository(ctrl *gomock.Controller) *MockQuoteRepository {
	mock := &MockQuoteRepository{ctrl: ctrl}
	mock.recorder = &MockQuoteRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuoteRepository) EXPECT() *MockQuoteRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockQuoteRepository) Create(ctx *appcontext.AppContext, quote domain.Quote) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, quote)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockQuoteRepositoryMockRecorder) Create(ctx, quote any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockQuoteRepository)(nil).Create), ctx, quote)
}

// FindLatest mocks base method.
func (m *MockQuoteRepository) FindLatest(ctx *appcontext.AppContext) (*domain.Quote, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindLatest", ctx)
	ret0, _ := ret[0].(*domain.Quote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindLatest indicates an expected call of FindLatest.
func (mr *MockQuoteRepositoryMockRecorder) FindLatest(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindLatest", reflect.TypeOf((*MockQuoteRepository)(nil).FindLatest), ctx)
}