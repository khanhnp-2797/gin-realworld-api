package tests

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
)unc setupArticleServiceTest(t *testing.T) (*mocks.MockArticleRepository, *mocks.MockUserRepository, services.ArticleService) {
	helpers.InitTestConfig()
func setupArticleServiceTest(t *testing.T) (*mocks.MockArticleRepository, *mocks.MockUserRepository, services.ArticleService) {
	helpers.InitTestConfig()wMockUserRepository()
	mockArticleRepo := mocks.NewMockArticleRepository()po, mockUserRepo)
	mockUserRepo := mocks.NewMockUserRepository()
	service := services.NewArticleService(mockArticleRepo, mockUserRepo)
	return mockArticleRepo, mockUserRepo, service
}unc TestArticleService_CreateArticle_Success(t *testing.T) {
	// Setup
func TestArticleService_CreateArticle_Success(t *testing.T) {eTest(t)
	// Setup
	mockArticleRepo, mockUserRepo, service := setupArticleServiceTest(t)D: 1, Name: "golang"}, nil)
	mockArticleRepo.On("FindOrCreateTag", "testing").Return(&models.Tag{ID: 2, Name: "testing"}, nil)
	mockArticleRepo.On("FindOrCreateTag", "golang").Return(&models.Tag{ID: 1, Name: "golang"}, nil)
	mockArticleRepo.On("FindOrCreateTag", "testing").Return(&models.Tag{ID: 2, Name: "testing"}, nil)
	mockArticleRepo.On("Create", mock.AnythingOfType("*models.Article")).Return(nil)
		ID:          1,
	mockArticle := &models.Article{
		ID:          1,est Article",
		Slug:        "test-article",on",
		Title:       "Test Article",
		Description: "Test Description",
		Body:        "Test Body",
		AuthorID:    1,
		Author: models.User{",
			ID:       1,
			Username: "testuser",g{{ID: 1, Name: "golang"}, {ID: 2, Name: "testing"}},
		},eatedAt: time.Now(),
		Tags:      []models.Tag{{ID: 1, Name: "golang"}, {ID: 2, Name: "testing"}},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}ockArticleRepo.On("FindBySlug", mock.AnythingOfType("string")).Return(mockArticle, nil)
	mockArticleRepo.On("IsFavorited", uint(1), uint(1)).Return(false, nil)
	mockArticleRepo.On("FindBySlug", mock.AnythingOfType("string")).Return(mockArticle, nil)
	mockArticleRepo.On("IsFavorited", uint(1), uint(1)).Return(false, nil)
	mockArticleRepo.On("GetFavoritesCount", uint(1)).Return(int64(0), nil)
	req.Article.Title = "Test Article"
	req := &dto.CreateArticleRequest{}scription"
	req.Article.Title = "Test Article"
	req.Article.Description = "Test Description"sting"}
	req.Article.Body = "Test Body"
	req.Article.TagList = []string{"golang", "testing"}
	response, err := service.CreateArticle(1, req)
	// Act
	response, err := service.CreateArticle(1, req)
	assert.NoError(t, err)
	// AsserttNil(t, response)
	assert.NoError(t, err)Article", response.Article.Title)
	assert.NotNil(t, response)icle.TagList, 2)
	assert.Equal(t, "Test Article", response.Article.Title)
	assert.Len(t, response.Article.TagList, 2)
	mockArticleRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}unc TestArticleService_GetArticle_Success(t *testing.T) {
	// Setup
func TestArticleService_GetArticle_Success(t *testing.T) {viceTest(t)
	// Setup
	mockArticleRepo, mockUserRepo, service := setupArticleServiceTest(t)
		ID:          1,
	mockArticle := &models.Article{
		ID:          1,est Article",
		Slug:        "test-article",on",
		Title:       "Test Article",
		Description: "Test Description",
		Body:        "Test Body",
		AuthorID:    1,
		Author: models.User{",
			ID:       1,
			Username: "testuser",g{{ID: 1, Name: "golang"}},
		},eatedAt: time.Now(),
		Tags:      []models.Tag{{ID: 1, Name: "golang"}},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}ockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	mockArticleRepo.On("IsFavorited", uint(1), uint(1)).Return(true, nil)
	mockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	mockArticleRepo.On("IsFavorited", uint(1), uint(1)).Return(true, nil)
	mockArticleRepo.On("GetFavoritesCount", uint(1)).Return(int64(5), nil)

	userID := uint(1)
	response, err := service.GetArticle("test-article", &userID)
	// Act
	response, err := service.GetArticle("test-article", &userID)
	assert.NoError(t, err)
	// Assertual(t, "test-article", response.Article.Slug)
	assert.NoError(t, err)ponse.Article.FavoritesCount)
	assert.Equal(t, "test-article", response.Article.Slug)
	assert.Equal(t, 5, response.Article.FavoritesCount)
	assert.True(t, response.Article.Favorited)
	mockArticleRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}unc TestArticleService_GetArticle_NotFound(t *testing.T) {
	// Setup
func TestArticleService_GetArticle_NotFound(t *testing.T) {iceTest(t)
	// Setup
	mockArticleRepo, mockUserRepo, service := setupArticleServiceTest(t)RecordNotFound)

	mockArticleRepo.On("FindBySlug", "nonexistent").Return(nil, gorm.ErrRecordNotFound)
	_, err := service.GetArticle("nonexistent", nil)
	// Act
	_, err := service.GetArticle("nonexistent", nil)
	assert.Error(t, err)
	// Assertual(t, "article not found", err.Error())
	assert.Error(t, err)rtExpectations(t)
	assert.Equal(t, "article not found", err.Error())
	mockArticleRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}unc TestArticleService_DeleteArticle_Success(t *testing.T) {
	// Setup
func TestArticleService_DeleteArticle_Success(t *testing.T) {eTest(t)
	// Setup
	mockArticleRepo, mockUserRepo, service := setupArticleServiceTest(t)
		ID:       1,
	mockArticle := &models.Article{
		ID:       1,
		Slug:     "test-article",
		AuthorID: 1,
	}ockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	mockArticleRepo.On("Delete", mockArticle).Return(nil)
	mockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	mockArticleRepo.On("Delete", mockArticle).Return(nil)
	err := service.DeleteArticle("test-article", 1)
	// Act
	err := service.DeleteArticle("test-article", 1)
	assert.NoError(t, err)
	// AssertleRepo.AssertExpectations(t)
	assert.NoError(t, err)ectations(t)
	mockArticleRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}unc TestArticleService_DeleteArticle_Unauthorized(t *testing.T) {
	// Setup
func TestArticleService_DeleteArticle_Unauthorized(t *testing.T) {(t)
	// Setup
	mockArticleRepo, mockUserRepo, service := setupArticleServiceTest(t)
		ID:       1,
	mockArticle := &models.Article{
		ID:       1,
		Slug:     "test-article",
		AuthorID: 1,
	}ockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)

	mockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	err := service.DeleteArticle("test-article", 2) // Different user
	// Act
	err := service.DeleteArticle("test-article", 2) // Different user
	assert.Error(t, err)
	// Assertual(t, "unauthorized to delete this article", err.Error())
	assert.Error(t, err)rtExpectations(t)
	assert.Equal(t, "unauthorized to delete this article", err.Error())
	mockArticleRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}unc TestArticleService_FavoriteArticle_Success(t *testing.T) {
	// Setup
func TestArticleService_FavoriteArticle_Success(t *testing.T) {est(t)
	// Setup
	mockArticleRepo, mockUserRepo, service := setupArticleServiceTest(t)
		ID:     1,
	mockArticle := &models.Article{
		ID:     1,dels.User{ID: 1, Username: "author"},
		Slug:   "test-article",
		Author: models.User{ID: 1, Username: "author"},
	}ockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	mockArticleRepo.On("FavoriteArticle", uint(2), uint(1)).Return(nil)
	mockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	mockArticleRepo.On("FavoriteArticle", uint(2), uint(1)).Return(nil)il)
	mockArticleRepo.On("IsFavorited", uint(2), uint(1)).Return(true, nil)
	mockArticleRepo.On("GetFavoritesCount", uint(1)).Return(int64(1), nil)
	response, err := service.FavoriteArticle("test-article", 2)
	// Act
	response, err := service.FavoriteArticle("test-article", 2)
	assert.NoError(t, err)
	// Assertue(t, response.Article.Favorited)
	assert.NoError(t, err)ponse.Article.FavoritesCount)
	assert.True(t, response.Article.Favorited)
	assert.Equal(t, 1, response.Article.FavoritesCount)
	mockArticleRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)}