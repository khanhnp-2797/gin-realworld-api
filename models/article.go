package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Slug        string         `gorm:"uniqueIndex;not null" json:"slug"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `gorm:"not null" json:"description"`
	Body        string         `gorm:"type:text;not null" json:"body"`
	AuthorID    uint           `gorm:"not null" json:"author_id"`
	Author      User           `gorm:"foreignKey:AuthorID" json:"author"`
	Tags        []Tag          `gorm:"many2many:article_tags;" json:"tags"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Tag struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"uniqueIndex;not null" json:"name"`
	Articles  []Article `gorm:"many2many:article_tags;" json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
