package internal

import (
	"fmt"
	"gorm.io/gorm"
	"main/internal/data/model"
	"os"
)

var DB *gorm.DB

func SetDB(db *gorm.DB) {
	DB = db
	err := DB.AutoMigrate(&model.UserProfile{})

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
