// Code generated by MockGen. DO NOT EDIT.
// Source: i_repository.go

// Package article is a generated GoMock package.
package article

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockRepository) Add(article Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", article)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockRepositoryMockRecorder) Add(article interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRepository)(nil).Add), article)
}

// GetList mocks base method.
func (m *MockRepository) GetList(keyword, author string) ([]Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", keyword, author)
	ret0, _ := ret[0].([]Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockRepositoryMockRecorder) GetList(keyword, author interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockRepository)(nil).GetList), keyword, author)
}
