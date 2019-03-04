// Code generated by MockGen. DO NOT EDIT.
// Source: article.go

// Package service is a generated GoMock package.
package service

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockArticleService is a mock of ArticleService interface
type MockArticleService struct {
	ctrl     *gomock.Controller
	recorder *MockArticleServiceMockRecorder
}

// MockArticleServiceMockRecorder is the mock recorder for MockArticleService
type MockArticleServiceMockRecorder struct {
	mock *MockArticleService
}

// NewMockArticleService creates a new mock instance
func NewMockArticleService(ctrl *gomock.Controller) *MockArticleService {
	mock := &MockArticleService{ctrl: ctrl}
	mock.recorder = &MockArticleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockArticleService) EXPECT() *MockArticleServiceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockArticleService) Create(name, content string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockArticleServiceMockRecorder) Create(name, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArticleService)(nil).Create), name, content)
}

// List mocks base method
func (m *MockArticleService) List() ([]ArticleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]ArticleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockArticleServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockArticleService)(nil).List))
}

// Like mocks base method
func (m *MockArticleService) Like(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Like", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Like indicates an expected call of Like
func (mr *MockArticleServiceMockRecorder) Like(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Like", reflect.TypeOf((*MockArticleService)(nil).Like), id)
}
