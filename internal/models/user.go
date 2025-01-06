package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string `gorm:"size:100;not null"`
	Email         string `gorm:"uniqueIndex;not null"`
	Password      string `gorm:"size:255"`
	OAuthProvider string `gorm:"size:50"`
	OAuthID       string `gorm:"size:100"`
}
