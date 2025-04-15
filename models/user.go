package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null" validate:"required,min=3,max=20"`
	Password string `json:"password" gorm:"not null" validate:"required,min=6"`
}
