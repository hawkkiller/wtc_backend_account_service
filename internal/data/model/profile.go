package model

import (
	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model `json:"-"`
	Username   string `json:"username" validate:"required" gorm:"unique"`
	Password   string `json:"password" validate:"required"`
	Email      string `json:"email" validate:"required" gorm:"index:unique"`
	Sex        string `json:"sex" validate:"required" gorm:"default:undefined"`
}
