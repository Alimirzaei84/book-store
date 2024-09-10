package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var(
	db *gorm.DB
)

func GetDB() *gorm.DB {
	return db
}

func Connect() {
	d, err := gorm.Open("mysql", "root:Mirzaeiali17965@tcp(localhost:3306)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local")

	// d, err := gorm.Open("mysql", "root:Mirzaeiali17965@12@/mydatabase?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}