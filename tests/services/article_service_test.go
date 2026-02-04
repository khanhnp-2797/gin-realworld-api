package services_test

import (
	"testing"
	"time"

	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"github.com/khanhnp-2797/gin-realworld-api/services"
	"github.com/khanhnp-2797/gin-realworld-api/tests/helpers"
	"github.com/khanhnp-2797/gin-realworld-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func setupArticleServiceTest(t *testing.T) (*mocks.MockArticleRepository, *mocks.MockUserRepository, services.ArticleService) {
	helpers.InitTestConfig()
	mockArticleRepo := mocks.NewMockArticleRepository()
	mockUserRepo := mocks.NewMockUserRepository()
	service := services.NewArticleService(mockArticleRepo, mockUserRepo)
	return mockArticleRepo, mockUserRepo, service
}

func TestArticleService_CreateArticle_Success(t *testing.T) {
	// Setup
	mockArticleRepo, mockUserRepo, service := setupArticleServiceTest(t)

	mockArticleRepo.On("FindOrCreateTag", "golang").Return(&models.Tag{ID: 1, Name: "golang"}, nil)
	mockArticleRepo.On("FindOrCreateTag", "testing").Return(&models.Tag{ID: 2, Name: "testing"}, nil)
	mockArticleRepo.On("Create", mock.AnythingOfType("*models.Article")).Return(nil)

	mockArticle := &models.Article{
		ID:          1,
		Slug:        "test-article",
		Title:       "Test Article",
		Description: "Test Description",
		Body:        "Test Body",
		AuthorID:    1,
		Author: models.User{
			ID:       1,
			Username: "testuser",
		},
		Tags:      []models.Tag{{ID: 1, Name: "golang"}, {ID: 2, Name: "testing"}},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockArticleRepo.On("FindBySlug", mock.AnythingOfType("string")).Return(mockArticle, nil)
	mockArticleRepo.On("IsFavorited", uint(1), uint(1)).Return(false, nil)
	mockArticleRepo.On("GetFavoritesCount", uint(1)).Return(int64(0), nil)

	req := &dto.CreateArticleRequest{}
	req.Article.Title = "Test Article"
	req.Article.Description = "Test Description"
	req.Article.Body = "Test Body"
	req.Article.TagList = []string{"golang", "testing"}

	// Act
	response, err := service.CreateArticle(1, req)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Test Article", response.Article.Title)
	assert.Len(t, response.Article.TagList, 2)
	mockArticleRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestArticleService_GetArticle_Success(t *testing.T) {
	// Setup
	mockArticleRepo, mockUserRepo, service := setupArticleServiceTest(t)

	mockArticle := &models.Article{
		ID:          1,
		Slug:        "test-article",
		Title:       "Test Article",
		Description: "Test Description",
		Body:        "Test Body",
		AuthorID:    1,
		Author: models.User{
			ID:       1,
			Username: "testuser",
		},
		Tags:      []models.Tag{{ID: 1, Name: "golang"}},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	mockArticleRepo.On("IsFavorited", uint(1), uint(1)).Return(true, nil)
	mockArticleRepo.On("GetFavoritesCount", uint(1)).Return(int64(5), nil)

	userID := uint(1)

	// Act
	response, err := service.GetArticle("test-article", &userID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "test-article", response.Article.Slug)
	assert.Equal(t, 5, response.Article.FavoritesCount)
	assert.True(t, response.Article.Favorited)
	mockArticleRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestArticleService_GetArticle_NotFound(t *testing.T) {
	// Setup
	mockArticleRepo, mockUserRepo, service := setupArticleServiceTest(t)

	mockArticleRepo.On("FindBySlug", "nonexistent").Return(nil, gorm.ErrRecordNotFound)

	// Act
	_, err := service.GetArticle("nonexistent", nil)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "article not found", err.Error())
	mockArticleRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestArticleService_DeleteArticle_Success(t *testing.T) {
	// Setup
	mockArticleRepo, mockUserRepo, service := setupArticleServiceTest(t)

	mockArticle := &models.Article{
		ID:       1,
		Slug:     "test-article",
		AuthorID: 1,
	}

	mockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	mockArticleRepo.On("Delete", mockArticle).Return(nil)

	// Act
	err := service.DeleteArticle("test-article", 1)

	// Assert
	assert.NoError(t, err)
	mockArticleRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestArticleService_DeleteArticle_Unauthorized(t *testing.T) {
	// Setup
	mockArticleRepo, mockUserRepo, service := setupArticleServiceTest(t)

	mockArticle := &models.Article{
		ID:       1,
		Slug:     "test-article",
		AuthorID: 1,
	}

	mockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)

	// Act
	err := service.DeleteArticle("test-article", 2) // Different user

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "unauthorized to delete this article", err.Error())
	mockArticleRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

func TestArticleService_FavoriteArticle_Success(t *testing.T) {
	// Setup
	mockArticleRepo, mockUserRepo, service := setupArticleServiceTest(t)

	mockArticle := &models.Article{
		ID:     1,
		Slug:   "test-article",
		Author: models.User{ID: 1, Username: "author"},
	}

	mockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	mockArticleRepo.On("FavoriteArticle", uint(2), uint(1)).Return(nil)
	mockArticleRepo.On("IsFavorited", uint(2), uint(1)).Return(true, nil)
	mockArticleRepo.On("GetFavoritesCount", uint(1)).Return(int64(1), nil)

	// Act
	response, err := service.FavoriteArticle("test-article", 2)

	// Assert
	assert.NoError(t, err)
	assert.True(t, response.Article.Favorited)
	assert.Equal(t, 1, response.Article.FavoritesCount)
	mockArticleRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}
