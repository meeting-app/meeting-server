package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite
)

// DB ...
var DB *gorm.DB

// Init init database
func Init() {
	db, err := gorm.Open("sqlite3", "meeting.db")
	if err != nil {
		panic(err)
	}
	DB = db
}

// GetDB get database instance
func GetDB() *gorm.DB {
	return DB
}
