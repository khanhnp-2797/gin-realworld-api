package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khanhnp-2797/gin-realworld-api/dto"
	"github.com/khanhnp-2797/gin-realworld-api/services"
	"github.com/khanhnp-2797/gin-realworld-api/utils"
)

type CommentController struct {
	commentService services.CommentService
}

func NewCommentController(commentService services.CommentService) *CommentController {
	return &CommentController{
		commentService: commentService,
	}
}

// AddComment - Thêm comment vào bài viết
// POST /api/articles/:slug/comments
func (cc *CommentController) AddComment(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.AuthError(""))
		return
	}

	slug := c.Param("slug")

	var req dto.AddCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ValidationError("Invalid request body"))
		return
	}

	response, err := cc.commentService.AddComment(slug, userID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetComments - Lấy danh sách comments của bài viết
// GET /api/articles/:slug/comments
func (cc *CommentController) GetComments(c *gin.Context) {
	slug := c.Param("slug")

	var userID *uint
	if id, exists := c.Get("userID"); exists {
		uid := id.(uint)
		userID = &uid
	}

	response, err := cc.commentService.GetComments(slug, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteComment - Xoá comment
// DELETE /api/articles/:slug/comments/:id
func (cc *CommentController) DeleteComment(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.AuthError(""))
		return
	}

	slug := c.Param("slug")
	commentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ValidationError("Invalid comment ID"))
		return
	}

	if err := cc.commentService.DeleteComment(slug, uint(commentID), userID.(uint)); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
