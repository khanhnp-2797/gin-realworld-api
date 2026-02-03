package dto

import "time"

// AddCommentRequest - Request thêm comment
type AddCommentRequest struct {
	Comment struct {
		Body string `json:"body" binding:"required,min=1,max=2000"`
	} `json:"comment" binding:"required"`
}

// CommentAuthor - Thông tin tác giả comment
type CommentAuthor struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

// CommentData - Dữ liệu comment
type CommentData struct {
	ID        uint          `json:"id"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	Body      string        `json:"body"`
	Author    CommentAuthor `json:"author"`
}

// CommentResponse - Response cho một comment
type CommentResponse struct {
	Comment CommentData `json:"comment"`
}

// CommentsResponse - Response cho danh sách comments
type CommentsResponse struct {
	Comments []CommentData `json:"comments"`
}
