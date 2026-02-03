package dto

import "time"

// CreateArticleRequest - Request tạo bài viết mới
type CreateArticleRequest struct {
	Article struct {
		Title       string   `json:"title" binding:"required,min=1,max=200"`
		Description string   `json:"description" binding:"required,min=1,max=500"`
		Body        string   `json:"body" binding:"required,min=1"`
		TagList     []string `json:"tagList"`
	} `json:"article" binding:"required"`
}

// UpdateArticleRequest - Request cập nhật bài viết
type UpdateArticleRequest struct {
	Article struct {
		Title       string `json:"title" binding:"omitempty,min=1,max=200"`
		Description string `json:"description" binding:"omitempty,min=1,max=500"`
		Body        string `json:"body" binding:"omitempty,min=1"`
	} `json:"article" binding:"required"`
}

// ArticleAuthor - Thông tin tác giả bài viết
type ArticleAuthor struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

// ArticleData - Dữ liệu bài viết
type ArticleData struct {
	Slug           string        `json:"slug"`
	Title          string        `json:"title"`
	Description    string        `json:"description"`
	Body           string        `json:"body"`
	TagList        []string      `json:"tagList"`
	CreatedAt      time.Time     `json:"createdAt"`
	UpdatedAt      time.Time     `json:"updatedAt"`
	Favorited      bool          `json:"favorited"`
	FavoritesCount int           `json:"favoritesCount"`
	Author         ArticleAuthor `json:"author"`
}

// ArticleResponse - Response cho một bài viết
type ArticleResponse struct {
	Article ArticleData `json:"article"`
}

// ArticlesResponse - Response cho danh sách bài viết
type ArticlesResponse struct {
	Articles      []ArticleData `json:"articles"`
	ArticlesCount int           `json:"articlesCount"`
}
