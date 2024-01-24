package dao

import (
	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	webfinger "github.com/Hana-ame/fedi-antenna/webfinger/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db, _ = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	activitypub.AutoMigrate(db)
	core.AutoMigrate(db)
	webfinger.AutoMigrate(db)
	entities.AutoMigrate(db)

}

func Where(query any, args ...any) (tx *gorm.DB) {
	return db.Where(query, args...)
}

func Create(o any) error {
	tx := db.Create(o)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

//	func Read(o any) error {
//		tx := db.First(o)
//		if tx.Error != nil {
//			return tx.Error
//		}
//		return nil
//	}
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
