package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"notNull;unique" json:"username"`
	Password  string `gorm:"notNull" json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Posts     []Post `json:"posts"`
}
