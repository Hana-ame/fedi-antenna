package db

import (
	"database/sql"

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

func DB() *gorm.DB {
	return db
}
func Begin(opts ...*sql.TxOptions) *gorm.DB {
	return db.Begin(opts...)
}
func AutoMigrate(dst ...any) error {
	return db.AutoMigrate(dst...)
}

func Create(tx *gorm.DB, o any) error {
	tx.Create(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func Read(tx *gorm.DB, o any) error {
	tx.Where(o).First(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func Update(tx *gorm.DB, o any) error {
	tx.Save(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func Delete(tx *gorm.DB, o any) error {
	tx.Delete(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
