package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null; size:60; unique"`
	Password string `json:"password" gorm:"not null"`
}
