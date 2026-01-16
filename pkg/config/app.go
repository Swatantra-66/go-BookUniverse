package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	dsn := os.Getenv("DSN")

	if dsn == "" {
		dsn = "root:user15@tcp(localhost:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"
	}

	database, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}
