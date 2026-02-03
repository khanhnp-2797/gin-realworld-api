package repositories

import (
	"github.com/khanhnp-2797/gin-realworld-api/models"
	"gorm.io/gorm"
)

type TagRepository interface {
	GetAllTags() ([]string, error)
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) GetAllTags() ([]string, error) {
	var tags []models.Tag
	err := r.db.Order("name ASC").Find(&tags).Error
	if err != nil {
		return nil, err
	}

	tagNames := make([]string, len(tags))
	for i, tag := range tags {
		tagNames[i] = tag.Name
	}
	return tagNames, nil
}
