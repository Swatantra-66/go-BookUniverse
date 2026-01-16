package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	dsn := "root:user15@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"
	database, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}
