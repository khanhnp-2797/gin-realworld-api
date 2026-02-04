package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khanhnp-2797/gin-realworld-api/controllers"
	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"github.com/khanhnp-2797/gin-realworld-api/services"
	"github.com/khanhnp-2797/gin-realworld-api/tests/helpers"
	"github.com/khanhnp-2797/gin-realworld-api/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestArticleController_CreateArticle_Success(t *testing.T) {
	// Setup
	helpers.InitTestConfig()
	mockArticleRepo := mocks.NewMockArticleRepository()
	mockUserRepo := mocks.NewMockUserRepository()
	service := services.NewArticleService(mockArticleRepo, mockUserRepo)
	controller := controllers.NewArticleController(service)

	mockArticleRepo.On("FindOrCreateTag", "golang").Return(&models.Tag{ID: 1, Name: "golang"}, nil)
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
		Tags:      []models.Tag{{ID: 1, Name: "golang"}},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockArticleRepo.On("FindBySlug", mock.AnythingOfType("string")).Return(mockArticle, nil)
	mockArticleRepo.On("IsFavorited", uint(1), uint(1)).Return(false, nil)
	mockArticleRepo.On("GetFavoritesCount", uint(1)).Return(int64(0), nil)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/api/articles", func(c *gin.Context) {
		c.Set("userID", uint(1))
		controller.CreateArticle(c)
	})

	createReq := dto.CreateArticleRequest{}
	createReq.Article.Title = "Test Article"
	createReq.Article.Description = "Test Description"
	createReq.Article.Body = "Test Body"
	createReq.Article.TagList = []string{"golang"}

	body, _ := json.Marshal(createReq)
	req := httptest.NewRequest("POST", "/api/articles", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)

	var response dto.ArticleResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "Test Article", response.Article.Title)
	mockArticleRepo.AssertExpectations(t)
}

func TestArticleController_GetArticle_Success(t *testing.T) {
	helpers.InitTestConfig()
	mockArticleRepo := mocks.NewMockArticleRepository()
	mockUserRepo := mocks.NewMockUserRepository()
	service := services.NewArticleService(mockArticleRepo, mockUserRepo)
	controller := controllers.NewArticleController(service)

	mockArticle := &models.Article{
		ID:          1,
		Slug:        "test-article",
		Title:       "Test Article",
		Description: "Test Description",
		Body:        "Test Body",
		Author: models.User{
			ID:       1,
			Username: "testuser",
		},
		Tags:      []models.Tag{{ID: 1, Name: "golang"}},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	mockArticleRepo.On("IsFavorited", uint(0), uint(1)).Return(false, nil)
	mockArticleRepo.On("GetFavoritesCount", uint(1)).Return(int64(5), nil)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/articles/:slug", controller.GetArticle)

	req := httptest.NewRequest("GET", "/api/articles/test-article", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response dto.ArticleResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "test-article", response.Article.Slug)
	assert.Equal(t, 5, response.Article.FavoritesCount)
}

func TestArticleController_DeleteArticle_Success(t *testing.T) {
	helpers.InitTestConfig()
	mockArticleRepo := mocks.NewMockArticleRepository()
	mockUserRepo := mocks.NewMockUserRepository()
	service := services.NewArticleService(mockArticleRepo, mockUserRepo)
	controller := controllers.NewArticleController(service)

	mockArticle := &models.Article{
		ID:       1,
		Slug:     "test-article",
		AuthorID: 1,
	}

	mockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	mockArticleRepo.On("Delete", mockArticle).Return(nil)

	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.DELETE("/api/articles/:slug", func(c *gin.Context) {
		c.Set("userID", uint(1))
		controller.DeleteArticle(c)
	})

	req := httptest.NewRequest("DELETE", "/api/articles/test-article", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockArticleRepo.AssertExpectations(t)
}
