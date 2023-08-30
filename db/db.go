package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db = initDB()
)

// init a db
func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func init() {
	Init()
}

func Init() {
	// Migrate the schema
	db.AutoMigrate(&Log{})
	db.AutoMigrate(&Object{})
}
