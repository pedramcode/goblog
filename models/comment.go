package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID  uint   `gorm:"notNull" json:"post_id"`
	Title   string `gorm:"notNull" json:"title"`
	Content string `gorm:"notNull" json:"content"`
}
