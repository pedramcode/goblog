package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID  uint   `gorm:"notNull" json:"user_id"`
	PostID  uint   `gorm:"notNull" json:"post_id"`
	Content string `gorm:"notNull" json:"content"`
}
