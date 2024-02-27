package entities

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/mastodon/entities/account"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities/status"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) (err error) {

	err = db.AutoMigrate(new(account.Field))
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(new(CustomEmoji))
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(new(MediaAttachment))
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(new(status.Tag))
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(new(CustomEmoji))
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(new(Account))
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(new(Status))
	if err != nil {
		log.Println(err)
	}
	return nil
}
