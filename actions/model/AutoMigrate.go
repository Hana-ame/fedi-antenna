package model

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&IDType{})
	db.AutoMigrate(&Image{})
	db.AutoMigrate(&PublicKey{})
	db.AutoMigrate(&Follow{})
	db.AutoMigrate(&Block{})
	db.AutoMigrate(&PublicKey{})
	db.AutoMigrate(&Undo{})
	db.AutoMigrate(&Accept{})
	db.AutoMigrate(&Reject{})
	db.AutoMigrate(&Note{})
}
