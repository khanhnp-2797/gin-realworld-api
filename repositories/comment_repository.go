package repositories

import (
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment *models.Comment) error
	FindByID(id uint) (*models.Comment, error)
	FindByArticleSlug(slug string) ([]models.Comment, error)
	Delete(comment *models.Comment) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) FindByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := r.db.Preload("Author").First(&comment, id).Error
	return &comment, err
}

func (r *commentRepository) FindByArticleSlug(slug string) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.db.
		Joins("JOIN articles ON articles.id = comments.article_id").
		Where("articles.slug = ?", slug).
		Preload("Author").
		Order("comments.created_at DESC").
		Find(&comments).Error
	return comments, err
}

func (r *commentRepository) Delete(comment *models.Comment) error {
	return r.db.Delete(comment).Error
}
