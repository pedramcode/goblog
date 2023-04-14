package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID   uint      `json:"user_id"`
	Title    string    `gorm:"notNull" json:"title"`
	Content  string    `gorm:"notNull" json:"content"`
	Comments []Comment `json:"comments"`
}
