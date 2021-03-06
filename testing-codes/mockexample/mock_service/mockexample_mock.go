// Code generated by MockGen. DO NOT EDIT.
// Source: mockexample.go

// Package mock_mockexample is a generated GoMock package.
package mock_service

import (
	gomock "github.com/golang/mock/gomock"
	mockexample "github.com/hgsgtk/go-snippets/testing-codes/mockexample"
	reflect "reflect"
)

// MockCustomerRepository is a mock of CustomerRepository interface
type MockCustomerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerRepositoryMockRecorder
}

// MockCustomerRepositoryMockRecorder is the mock recorder for MockCustomerRepository
type MockCustomerRepositoryMockRecorder struct {
	mock *MockCustomerRepository
}

// NewMockCustomerRepository creates a new mock instance
func NewMockCustomerRepository(ctrl *gomock.Controller) *MockCustomerRepository {
	mock := &MockCustomerRepository{ctrl: ctrl}
	mock.recorder = &MockCustomerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCustomerRepository) EXPECT() *MockCustomerRepositoryMockRecorder {
	return m.recorder
}

// GetByCode mocks base method
func (m *MockCustomerRepository) GetByCode(code string) (mockexample.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCode", code)
	ret0, _ := ret[0].(mockexample.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCode indicates an expected call of GetByCode
func (mr *MockCustomerRepositoryMockRecorder) GetByCode(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCode", reflect.TypeOf((*MockCustomerRepository)(nil).GetByCode), code)
}

// MockTripRepository is a mock of TripRepository interface
type MockTripRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTripRepositoryMockRecorder
}

// MockTripRepositoryMockRecorder is the mock recorder for MockTripRepository
type MockTripRepositoryMockRecorder struct {
	mock *MockTripRepository
}

// NewMockTripRepository creates a new mock instance
func NewMockTripRepository(ctrl *gomock.Controller) *MockTripRepository {
	mock := &MockTripRepository{ctrl: ctrl}
	mock.recorder = &MockTripRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTripRepository) EXPECT() *MockTripRepositoryMockRecorder {
	return m.recorder
}

// GetByCustomerID mocks base method
func (m *MockTripRepository) GetByCustomerID(id int, status string) (mockexample.Trip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCustomerID", id, status)
	ret0, _ := ret[0].(mockexample.Trip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCustomerID indicates an expected call of GetByCustomerID
func (mr *MockTripRepositoryMockRecorder) GetByCustomerID(id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCustomerID", reflect.TypeOf((*MockTripRepository)(nil).GetByCustomerID), id, status)
}
