package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/services"
	"github.com/khanhnp-2797/gin-realworld-api/utils"
)

type ArticleController struct {
	articleService services.ArticleService
}

func NewArticleController(articleService services.ArticleService) *ArticleController {
	return &ArticleController{
		articleService: articleService,
	}
}

// CreateArticle - Tạo bài viết mới
// POST /api/articles
func (ac *ArticleController) CreateArticle(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.AuthError(""))
		return
	}

	var req dto.CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ValidationError("Invalid request body"))
		return
	}

	response, err := ac.articleService.CreateArticle(userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetArticle - Lấy chi tiết một bài viết
// GET /api/articles/:slug
func (ac *ArticleController) GetArticle(c *gin.Context) {
	slug := c.Param("slug")

	var userID *uint
	if id, exists := c.Get("userID"); exists {
		uid := id.(uint)
		userID = &uid
	}

	response, err := ac.articleService.GetArticle(slug, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.NotFoundError("Article"))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateArticle - Cập nhật bài viết
// PUT /api/articles/:slug
func (ac *ArticleController) UpdateArticle(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.AuthError(""))
		return
	}

	slug := c.Param("slug")

	var req dto.UpdateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ValidationError("Invalid request body"))
		return
	}

	response, err := ac.articleService.UpdateArticle(slug, userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteArticle - Xoá bài viết
// DELETE /api/articles/:slug
func (ac *ArticleController) DeleteArticle(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.AuthError(""))
		return
	}

	slug := c.Param("slug")

	if err := ac.articleService.DeleteArticle(slug, userID.(uint)); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetFeed - Lấy feed bài viết
// GET /api/articles/feed
func (ac *ArticleController) GetFeed(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.AuthError(""))
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	response, err := ac.articleService.GetFeed(userID.(uint), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}
