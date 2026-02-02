package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Body      string         `gorm:"type:text;not null" json:"body"`
	ArticleID uint           `gorm:"not null;index" json:"article_id"`
	Article   Article        `gorm:"foreignKey:ArticleID" json:"article"`
	AuthorID  uint           `gorm:"not null;index" json:"author_id"`
	Author    User           `gorm:"foreignKey:AuthorID" json:"author"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
