package dao

import (
	"log"

	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db, _ = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

func init() {
	db.AutoMigrate(&activitypub.User{})
	db.AutoMigrate(&activitypub.Endpoints{})
	db.AutoMigrate(&activitypub.IDType{})
	db.AutoMigrate(&activitypub.Image{})
	db.AutoMigrate(&activitypub.PublicKey{})
}

func Test() {
	log.Println("")
}

func Create(user any) error {
	tx := db.Create(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func Read(user any) ( error) {
	tx := db.Take(user)
	if tx.Error != nil {
		return  tx.Error
	}
	return  nil
}
func Update(user any) error {
	tx := db.Save(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
