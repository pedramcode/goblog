package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID   uint
	Title    string `gorm:"notNull"`
	Content  string `gorm:"notNull"`
	Comments []Comment
}
