package repositories

import (
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(article *models.Article) error
	FindBySlug(slug string) (*models.Article, error)
	Update(article *models.Article) error
	Delete(article *models.Article) error
	GetFeed(userID uint, limit, offset int) ([]models.Article, error)
	FindOrCreateTag(tagName string) (*models.Tag, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) Create(article *models.Article) error {
	return r.db.Create(article).Error
}

func (r *articleRepository) FindBySlug(slug string) (*models.Article, error) {
	var article models.Article
	err := r.db.Preload("Author").Preload("Tags").Where("slug = ?", slug).First(&article).Error
	return &article, err
}

func (r *articleRepository) Update(article *models.Article) error {
	return r.db.Save(article).Error
}

func (r *articleRepository) Delete(article *models.Article) error {
	return r.db.Delete(article).Error
}

func (r *articleRepository) GetFeed(userID uint, limit, offset int) ([]models.Article, error) {
	var articles []models.Article
	err := r.db.
		Joins("JOIN follows ON follows.following_id = articles.author_id").
		Where("follows.follower_id = ?", userID).
		Preload("Author").
		Preload("Tags").
		Order("articles.created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&articles).Error
	return articles, err
}

func (r *articleRepository) FindOrCreateTag(tagName string) (*models.Tag, error) {
	var tag models.Tag
	err := r.db.Where("name = ?", tagName).FirstOrCreate(&tag, models.Tag{Name: tagName}).Error
	return &tag, err
}
