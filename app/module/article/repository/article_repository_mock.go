// Code generated by MockGen. DO NOT EDIT.
// Source: article_repository.go

// Package mock_repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	ent "github.com/BerlitzPlatina/gf-uma/internal/ent"
	request "github.com/BerlitzPlatina/gf-uma/app/module/article/request"
	gomock "github.com/golang/mock/gomock"
)

// MockIArticleRepository is a mock of IArticleRepository interface.
type MockIArticleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIArticleRepositoryMockRecorder
}

// MockIArticleRepositoryMockRecorder is the mock recorder for MockIArticleRepository.
type MockIArticleRepositoryMockRecorder struct {
	mock *MockIArticleRepository
}

// NewMockIArticleRepository creates a new mock instance.
func NewMockIArticleRepository(ctrl *gomock.Controller) *MockIArticleRepository {
	mock := &MockIArticleRepository{ctrl: ctrl}
	mock.recorder = &MockIArticleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIArticleRepository) EXPECT() *MockIArticleRepositoryMockRecorder {
	return m.recorder
}

// CreateArticle mocks base method.
func (m *MockIArticleRepository) CreateArticle(request request.ArticleRequest) (*ent.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArticle", request)
	ret0, _ := ret[0].(*ent.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArticle indicates an expected call of CreateArticle.
func (mr *MockIArticleRepositoryMockRecorder) CreateArticle(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArticle", reflect.TypeOf((*MockIArticleRepository)(nil).CreateArticle), request)
}

// DeleteArticle mocks base method.
func (m *MockIArticleRepository) DeleteArticle(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArticle", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteArticle indicates an expected call of DeleteArticle.
func (mr *MockIArticleRepositoryMockRecorder) DeleteArticle(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArticle", reflect.TypeOf((*MockIArticleRepository)(nil).DeleteArticle), id)
}

// GetArticleByID mocks base method.
func (m *MockIArticleRepository) GetArticleByID(id int) (*ent.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticleByID", id)
	ret0, _ := ret[0].(*ent.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticleByID indicates an expected call of GetArticleByID.
func (mr *MockIArticleRepositoryMockRecorder) GetArticleByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticleByID", reflect.TypeOf((*MockIArticleRepository)(nil).GetArticleByID), id)
}

// GetArticles mocks base method.
func (m *MockIArticleRepository) GetArticles() ([]*ent.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticles")
	ret0, _ := ret[0].([]*ent.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticles indicates an expected call of GetArticles.
func (mr *MockIArticleRepositoryMockRecorder) GetArticles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticles", reflect.TypeOf((*MockIArticleRepository)(nil).GetArticles))
}

// UpdateArticle mocks base method.
func (m *MockIArticleRepository) UpdateArticle(id int, request request.ArticleRequest) (*ent.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateArticle", id, request)
	ret0, _ := ret[0].(*ent.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateArticle indicates an expected call of UpdateArticle.
func (mr *MockIArticleRepositoryMockRecorder) UpdateArticle(id, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArticle", reflect.TypeOf((*MockIArticleRepository)(nil).UpdateArticle), id, request)
}
