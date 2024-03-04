package dao

import (
	"github.com/Hana-ame/fedi-antenna/core/mydb"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

var DB *mydb.DB

func init() {
	var err error
	DB, err = mydb.NewDB("mastodon.db", entities.AutoMigrate)
	if err != nil {
		panic(err)
	}
	DB.DB().AutoMigrate(new(Relation))
	DB.DB().AutoMigrate(new(Reblog))
	DB.DB().AutoMigrate(new(Favourite))
}

const (
	RelationStatusPadding  = "padding"
	RelationStatusAccepted = "accepted"

	RelationStatusBlocking = "blocking"
)

const (
	VisibilityPublic = "public"
)

// actor follows object, activitypub one.
type Relation struct {
	ID int `gorm:"primarykey;autoIncrement"`

	Actor  string `gorm:"uniqueIndex:idx_actor_object"`
	Object string `gorm:"uniqueIndex:idx_actor_object"`

	// "padding" | "accepted" | "blocking"
	Status string
}

// all activitypub one.
type Reblog struct {
	Id     string `gorm:"primarykey"`
	Actor  string
	Object string

	Owner      string
	Visibility string

	DeletedAt int64
}

// all activitypub one.
type Favourite struct {
	ID     int    `gorm:"primarykey;autoIncrement"`
	Actor  string `gorm:"uniqueIndex:idx_actor_object"`
	Object string `gorm:"uniqueIndex:idx_actor_object"`

	Owner string

	DeletedAt int64
}
