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
	db.AutoMigrate(&activitypub.Follow{})
	db.AutoMigrate(&activitypub.Block{})
	db.AutoMigrate(&activitypub.PublicKey{})
}

func Test() {
	log.Println("")
}

func Create(o any) error {
	tx := db.Create(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func Read(o any) ( error) {
	tx := db.Take(o)
	if tx.Error != nil {
		return  tx.Error
	}
	return  nil
}
func Update(o any) error {
	tx := db.Save(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func Delete(o any) error{
	tx:= db.Delete(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
