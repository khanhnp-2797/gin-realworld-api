package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khanhnp-2797/gin-realworld-api/services"
	"github.com/khanhnp-2797/gin-realworld-api/utils"
)

type TagController struct {
	tagService services.TagService
}

func NewTagController(tagService services.TagService) *TagController {
	return &TagController{
		tagService: tagService,
	}
}

// GetTags - Lấy danh sách tags
// GET /api/tags
func (tc *TagController) GetTags(c *gin.Context) {
	response, err := tc.tagService.GetTags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response)
}
