package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open("sqlite3", "meeting.db")
	if err != nil {
		panic(err)
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
