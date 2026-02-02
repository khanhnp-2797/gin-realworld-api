package models

import "time"

type Follow struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	FollowerID  uint      `gorm:"not null;index" json:"follower_id"`
	FollowingID uint      `gorm:"not null;index" json:"following_id"`
	CreatedAt   time.Time `json:"created_at"`
}
