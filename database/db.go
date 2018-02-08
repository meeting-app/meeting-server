package database

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func Init() *gorm.DB {
	// TODO: Set params gorm.Open
	db, err := gorm.Open("", "")
	if err != nil {
		panic(err)
	}
	DB = db
	return DB
}

func getDB() *gorm.DB {
	return DB
}
