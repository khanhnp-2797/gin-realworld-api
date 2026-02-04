package mocks

import (
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"github.com/stretchr/testify/mock"
)

type MockCommentRepository struct {
	mock.Mock
}

func NewMockCommentRepository() *MockCommentRepository {
	return &MockCommentRepository{}
}

func (m *MockCommentRepository) Create(comment *models.Comment) error {
	args := m.Called(comment)
	return args.Error(0)
}

func (m *MockCommentRepository) FindByID(id uint) (*models.Comment, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Comment), args.Error(1)
}

func (m *MockCommentRepository) FindByArticleSlug(slug string) ([]models.Comment, error) {
	args := m.Called(slug)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Comment), args.Error(1)
}

func (m *MockCommentRepository) Delete(comment *models.Comment) error {
	args := m.Called(comment)
	return args.Error(0)
}
