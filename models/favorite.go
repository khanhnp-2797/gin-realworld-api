package models

import "time"

type Favorite struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	ArticleID uint      `gorm:"not null;index" json:"article_id"`
	CreatedAt time.Time `json:"created_at"`
}
