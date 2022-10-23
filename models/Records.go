package models

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	Record string `json:"record"`
	UserId int    `json:"user_id"`
	User   User   `json:"user" gorm:"foreignKey:user_id; References:id"`
}
