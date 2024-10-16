package dao

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(gormDB *gorm.DB) {
	db = gormDB
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.IDType{})
	db.AutoMigrate(&model.Image{})
	db.AutoMigrate(&model.PublicKey{})
	db.AutoMigrate(&model.Follow{})
	db.AutoMigrate(&model.Block{})
	db.AutoMigrate(&model.PublicKey{})
	db.AutoMigrate(&model.Undo{})
	db.AutoMigrate(&model.Accept{})
	db.AutoMigrate(&model.Reject{})
	db.AutoMigrate(&model.Note{})
}

// func Where(query any, args ...any) (tx *gorm.DB) {
// 	return db.Where(query, args...)
// }

func Create(o any) error {
	tx := db.Create(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func Read(o any) error {
	tx := db.Where(o).First(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func Update(o any) error {
	tx := db.Save(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func Delete(o any) error {
	tx := db.Delete(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DB() *gorm.DB {
	return db
}
