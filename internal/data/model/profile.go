package model

import (
	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model `json:"-"`
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Email      string `json:"email" validate:"required"`
	Sex        string `json:"sex" validate:"required"`
}
