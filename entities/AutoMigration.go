package entities

import (
	"github.com/Hana-ame/fedi-antenna/mastodon/entities/account"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities/status"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	db.AutoMigrate(new(account.Field))
	db.AutoMigrate(new(CustomEmoji))
	db.AutoMigrate(new(MediaAttachment))
	db.AutoMigrate(new(status.Tag))
	db.AutoMigrate(new(CustomEmoji))
	db.AutoMigrate(new(Account))
	db.AutoMigrate(new(Status))
	return nil
}
