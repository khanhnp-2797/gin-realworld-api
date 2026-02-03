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
	FavoriteArticle(userID, articleID uint) error
	UnfavoriteArticle(userID, articleID uint) error
	IsFavorited(userID, articleID uint) (bool, error)
	GetFavoritesCount(articleID uint) (int64, error)
	GetArticles(limit, offset int, tag, author, favorited string) ([]models.Article, error)
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

func (r *articleRepository) FavoriteArticle(userID, articleID uint) error {
	favorite := models.Favorite{
		UserID:    userID,
		ArticleID: articleID,
	}
	return r.db.FirstOrCreate(&favorite, favorite).Error
}

func (r *articleRepository) UnfavoriteArticle(userID, articleID uint) error {
	return r.db.Where("user_id = ? AND article_id = ?", userID, articleID).
		Delete(&models.Favorite{}).Error
}

func (r *articleRepository) IsFavorited(userID, articleID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Favorite{}).
		Where("user_id = ? AND article_id = ?", userID, articleID).
		Count(&count).Error
	return count > 0, err
}

func (r *articleRepository) GetFavoritesCount(articleID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Favorite{}).
		Where("article_id = ?", articleID).
		Count(&count).Error
	return count, err
}

func (r *articleRepository) GetArticles(limit, offset int, tag, author, favorited string) ([]models.Article, error) {
	var articles []models.Article
	query := r.db.Preload("Author").Preload("Tags")

	// Filter by tag
	if tag != "" {
		query = query.Joins("JOIN article_tags ON article_tags.article_id = articles.id").
			Joins("JOIN tags ON tags.id = article_tags.tag_id").
			Where("tags.name = ?", tag)
	}

	// Filter by author
	if author != "" {
		query = query.Joins("JOIN users ON users.id = articles.author_id").
			Where("users.username = ?", author)
	}

	// Filter by favorited
	if favorited != "" {
		query = query.Joins("JOIN favorites ON favorites.article_id = articles.id").
			Joins("JOIN users AS fav_users ON fav_users.id = favorites.user_id").
			Where("fav_users.username = ?", favorited)
	}

	err := query.Order("articles.created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&articles).Error

	return articles, err
}
