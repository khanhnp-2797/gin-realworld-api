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

func setupCommentServiceTest(t *testing.T) (*mocks.MockCommentRepository, *mocks.MockArticleRepository, services.CommentService) {
	helpers.InitTestConfig()
	mockCommentRepo := mocks.NewMockCommentRepository()
	mockArticleRepo := mocks.NewMockArticleRepository()
	service := services.NewCommentService(mockCommentRepo, mockArticleRepo)
	return mockCommentRepo, mockArticleRepo, service
}

func TestCommentService_AddComment_Success(t *testing.T) {
	// Setup
	mockCommentRepo, mockArticleRepo, service := setupCommentServiceTest(t)

	mockArticle := &models.Article{
		ID:   1,
		Slug: "test-article",
	}

	mockComment := &models.Comment{
		ID:        1,
		Body:      "Great article!",
		ArticleID: 1,
		AuthorID:  1,
		Author: models.User{
			ID:       1,
			Username: "commenter",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	mockCommentRepo.On("Create", mock.AnythingOfType("*models.Comment")).Return(nil)
	// Use mock.MatchedBy để accept bất kỳ uint nào (vì comment.ID có thể là 0 lúc tạo)
	mockCommentRepo.On("FindByID", mock.MatchedBy(func(id uint) bool {
		return id > 0 || id == 0 // Accept any ID
	})).Return(mockComment, nil)

	req := &dto.AddCommentRequest{}
	req.Comment.Body = "Great article!"

	// Act
	response, err := service.AddComment("test-article", 1, req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Great article!", response.Comment.Body)
	assert.Equal(t, "commenter", response.Comment.Author.Username)
	mockCommentRepo.AssertExpectations(t)
	mockArticleRepo.AssertExpectations(t)
}

func TestCommentService_AddComment_ArticleNotFound(t *testing.T) {
	mockCommentRepo, mockArticleRepo, service := setupCommentServiceTest(t)

	mockArticleRepo.On("FindBySlug", "nonexistent").Return(nil, gorm.ErrRecordNotFound)

	req := &dto.AddCommentRequest{}
	req.Comment.Body = "Great article!"

	// Act
	_, err := service.AddComment("nonexistent", 1, req)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "article not found", err.Error())
	mockCommentRepo.AssertExpectations(t)
	mockArticleRepo.AssertExpectations(t)
}

func TestCommentService_GetComments_Success(t *testing.T) {
	mockCommentRepo, mockArticleRepo, service := setupCommentServiceTest(t)

	mockArticle := &models.Article{
		ID:   1,
		Slug: "test-article",
	}

	mockCommentsList := []models.Comment{
		{
			ID:        1,
			Body:      "First comment",
			ArticleID: 1,
			AuthorID:  1,
			Author: models.User{
				ID:       1,
				Username: "user1",
			},
		},
		{
			ID:        2,
			Body:      "Second comment",
			ArticleID: 1,
			AuthorID:  2,
			Author: models.User{
				ID:       2,
				Username: "user2",
			},
		},
	}

	mockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	mockCommentRepo.On("FindByArticleSlug", "test-article").Return(mockCommentsList, nil)

	// Act
	response, err := service.GetComments("test-article", nil)

	// Assert
	assert.NoError(t, err)
	assert.Len(t, response.Comments, 2)
	assert.Equal(t, "First comment", response.Comments[0].Body)
	mockCommentRepo.AssertExpectations(t)
	mockArticleRepo.AssertExpectations(t)
}

func TestCommentService_DeleteComment_Success(t *testing.T) {
	mockCommentRepo, mockArticleRepo, service := setupCommentServiceTest(t)

	mockArticle := &models.Article{
		ID:   1,
		Slug: "test-article",
	}

	mockComment := &models.Comment{
		ID:        1,
		Body:      "Test comment",
		ArticleID: 1,
		AuthorID:  1,
	}

	mockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	mockCommentRepo.On("FindByID", uint(1)).Return(mockComment, nil)
	mockCommentRepo.On("Delete", mockComment).Return(nil)

	// Act
	err := service.DeleteComment("test-article", 1, 1)

	// Assert
	assert.NoError(t, err)
	mockCommentRepo.AssertExpectations(t)
	mockArticleRepo.AssertExpectations(t)
}

func TestCommentService_DeleteComment_Unauthorized(t *testing.T) {
	mockCommentRepo, mockArticleRepo, service := setupCommentServiceTest(t)

	mockArticle := &models.Article{
		ID:   1,
		Slug: "test-article",
	}

	mockComment := &models.Comment{
		ID:        1,
		Body:      "Test comment",
		ArticleID: 1,
		AuthorID:  1,
	}

	mockArticleRepo.On("FindBySlug", "test-article").Return(mockArticle, nil)
	mockCommentRepo.On("FindByID", uint(1)).Return(mockComment, nil)

	// Act
	err := service.DeleteComment("test-article", 1, 2)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "unauthorized to delete this comment", err.Error())
	mockCommentRepo.AssertExpectations(t)
	mockArticleRepo.AssertExpectations(t)
}
