// Code generated by MockGen. DO NOT EDIT.
// Source: userBalanceRepository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	reflect "reflect"

	models "github.com/dcwk/gophermart/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUserBalanceRepository is a mock of UserBalanceRepository interface.
type MockUserBalanceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserBalanceRepositoryMockRecorder
}

// MockUserBalanceRepositoryMockRecorder is the mock recorder for MockUserBalanceRepository.
type MockUserBalanceRepositoryMockRecorder struct {
	mock *MockUserBalanceRepository
}

// NewMockUserBalanceRepository creates a new mock instance.
func NewMockUserBalanceRepository(ctrl *gomock.Controller) *MockUserBalanceRepository {
	mock := &MockUserBalanceRepository{ctrl: ctrl}
	mock.recorder = &MockUserBalanceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserBalanceRepository) EXPECT() *MockUserBalanceRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserBalanceRepository) Create(ctx context.Context, userBalance *models.UserBalance) (*models.UserBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userBalance)
	ret0, _ := ret[0].(*models.UserBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserBalanceRepositoryMockRecorder) Create(ctx, userBalance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserBalanceRepository)(nil).Create), ctx, userBalance)
}

// GetUserBalanceByID mocks base method.
func (m *MockUserBalanceRepository) GetUserBalanceByID(ctx context.Context, userID int64, withLock bool) (*models.UserBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserBalanceByID", ctx, userID, withLock)
	ret0, _ := ret[0].(*models.UserBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserBalanceByID indicates an expected call of GetUserBalanceByID.
func (mr *MockUserBalanceRepositoryMockRecorder) GetUserBalanceByID(ctx, userID, withLock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserBalanceByID", reflect.TypeOf((*MockUserBalanceRepository)(nil).GetUserBalanceByID), ctx, userID, withLock)
}

// Update mocks base method.
func (m *MockUserBalanceRepository) Update(ctx context.Context, userBalance *models.UserBalance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, userBalance)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserBalanceRepositoryMockRecorder) Update(ctx, userBalance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserBalanceRepository)(nil).Update), ctx, userBalance)
}
