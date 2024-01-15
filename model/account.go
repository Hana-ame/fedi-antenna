package model

import (
	"github.com/Hana-ame/orderedmap"
	"gorm.io/gorm"
)

// only for actor's cache
type AccountID struct {
	Account         string                 `gorm:"primarykey"`
	WebfingerObject *orderedmap.OrderedMap `gorm:"serializer:json"`
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(new(AccountID))
}
