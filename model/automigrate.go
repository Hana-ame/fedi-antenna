package model

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(new(LocalUser))
	db.AutoMigrate(new(LocalNote))
	db.AutoMigrate(new(LocalNotify))
	db.AutoMigrate(new(LocalRelation))
}
