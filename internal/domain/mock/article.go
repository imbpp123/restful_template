// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/article.go
//
// Generated by this command:
//
//	mockgen -source internal/domain/article.go -destination internal/domain/mock/article.go -package=mockDomain
//

// Package mockDomain is a generated GoMock package.
package mockDomain

import (
	data "app/internal/data"
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockArticleRepository is a mock of ArticleRepository interface.
type MockArticleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockArticleRepositoryMockRecorder
	isgomock struct{}
}

// MockArticleRepositoryMockRecorder is the mock recorder for MockArticleRepository.
type MockArticleRepositoryMockRecorder struct {
	mock *MockArticleRepository
}

// NewMockArticleRepository creates a new mock instance.
func NewMockArticleRepository(ctrl *gomock.Controller) *MockArticleRepository {
	mock := &MockArticleRepository{ctrl: ctrl}
	mock.recorder = &MockArticleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleRepository) EXPECT() *MockArticleRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockArticleRepository) Create(ctx context.Context, article *data.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, article)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockArticleRepositoryMockRecorder) Create(ctx, article any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArticleRepository)(nil).Create), ctx, article)
}

// DeleteByUUID mocks base method.
func (m *MockArticleRepository) DeleteByUUID(ctx context.Context, uuid uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByUUID", ctx, uuid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByUUID indicates an expected call of DeleteByUUID.
func (mr *MockArticleRepositoryMockRecorder) DeleteByUUID(ctx, uuid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByUUID", reflect.TypeOf((*MockArticleRepository)(nil).DeleteByUUID), ctx, uuid)
}

// FindByUUID mocks base method.
func (m *MockArticleRepository) FindByUUID(ctx context.Context, uuid uuid.UUID) (*data.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUUID", ctx, uuid)
	ret0, _ := ret[0].(*data.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUUID indicates an expected call of FindByUUID.
func (mr *MockArticleRepositoryMockRecorder) FindByUUID(ctx, uuid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUUID", reflect.TypeOf((*MockArticleRepository)(nil).FindByUUID), ctx, uuid)
}

// List mocks base method.
func (m *MockArticleRepository) List(ctx context.Context, params *data.ArticleListParameters) ([]data.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, params)
	ret0, _ := ret[0].([]data.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockArticleRepositoryMockRecorder) List(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockArticleRepository)(nil).List), ctx, params)
}

// Update mocks base method.
func (m *MockArticleRepository) Update(ctx context.Context, article *data.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, article)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockArticleRepositoryMockRecorder) Update(ctx, article any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArticleRepository)(nil).Update), ctx, article)
}
