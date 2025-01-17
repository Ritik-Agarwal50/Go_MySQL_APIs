package config

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "ritik:root@tcp(localhost:3306)/Go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("database not found: ",err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
