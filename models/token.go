package models

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Key    string `gorm:"notNull" json:"key"`
}
