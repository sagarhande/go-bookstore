package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sagarhande/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	// Every DB neet to have init method
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}
