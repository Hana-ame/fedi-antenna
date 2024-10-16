package activitypub

import (
	"github.com/Hana-ame/fedi-antenna/activitypub/dao/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
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
