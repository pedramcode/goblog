package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostID  uint   `gorm:"notNull"`
	Title   string `gorm:"notNull"`
	Content string `gorm:"notNull"`
}
