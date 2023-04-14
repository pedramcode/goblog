package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"notNull;unique"`
	Password  string `gorm:"notNull"`
	Firstname string
	Lastname  string
	Email     string `gorm:"notNull;unique"`
	Phone     string
	Posts     []Post
}
