package mocks

import (
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"github.com/stretchr/testify/mock"
)

type MockArticleRepository struct {
	mock.Mock
}

func NewMockArticleRepository() *MockArticleRepository {
	return &MockArticleRepository{}
}

func (m *MockArticleRepository) Create(article *models.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func (m *MockArticleRepository) FindBySlug(slug string) (*models.Article, error) {
	args := m.Called(slug)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Article), args.Error(1)
}

func (m *MockArticleRepository) Update(article *models.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func (m *MockArticleRepository) Delete(article *models.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func (m *MockArticleRepository) GetFeed(userID uint, limit, offset int) ([]models.Article, error) {
	args := m.Called(userID, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Article), args.Error(1)
}

func (m *MockArticleRepository) FindOrCreateTag(tagName string) (*models.Tag, error) {
	args := m.Called(tagName)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Tag), args.Error(1)
}

func (m *MockArticleRepository) GetArticles(limit, offset int, tag, author, favorited string) ([]models.Article, error) {
	args := m.Called(limit, offset, tag, author, favorited)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Article), args.Error(1)
}

func (m *MockArticleRepository) FavoriteArticle(userID, articleID uint) error {
	args := m.Called(userID, articleID)
	return args.Error(0)
}

func (m *MockArticleRepository) UnfavoriteArticle(userID, articleID uint) error {
	args := m.Called(userID, articleID)
	return args.Error(0)
}

func (m *MockArticleRepository) IsFavorited(userID, articleID uint) (bool, error) {
	args := m.Called(userID, articleID)
	return args.Bool(0), args.Error(1)
}

func (m *MockArticleRepository) GetFavoritesCount(articleID uint) (int64, error) {
	args := m.Called(articleID)
	return args.Get(0).(int64), args.Error(1)
}
