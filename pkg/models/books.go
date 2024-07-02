package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ritik-agarwal50/go-bookstore/pkg/config"

)


var db *gorm.DB
type Book struct {
	gorm.Model
	Name string `gorn:""json:"name"`
	Author string `json:"author"`
	Publlication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}
