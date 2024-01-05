package dao

import (
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db, _ = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	db.AutoMigrate(&activitypub.User{})
	db.AutoMigrate(&activitypub.IDType{})
	db.AutoMigrate(&activitypub.Image{})
	db.AutoMigrate(&activitypub.PublicKey{})
	db.AutoMigrate(&activitypub.Follow{})
	db.AutoMigrate(&activitypub.Block{})
	db.AutoMigrate(&activitypub.PublicKey{})
	db.AutoMigrate(&activitypub.Undo{})
	db.AutoMigrate(&activitypub.Accept{})
	db.AutoMigrate(&activitypub.Reject{})
	db.AutoMigrate(&activitypub.Note{})

	db.AutoMigrate(&core.FediStatus{})

	core.AutoMigrate(db)
}

func Create(o any) error {
	tx := db.Create(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func Read(o any) error {
	tx := db.Take(o)
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
