package dao

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/db"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities/account"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities/status"
	"gorm.io/gorm"
)

func init() {
	err := db.AutoMigrate(new(MyRelation), new(MyReblog), new(MyFavourite))
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(new(account.Field))
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(new(entities.CustomEmoji))
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(new(entities.MediaAttachment))
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(new(status.Tag))
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(new(entities.CustomEmoji))
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(new(entities.Account))
	if err != nil {
		log.Println(err)
	}
	err = db.AutoMigrate(new(entities.Status))
	if err != nil {
		log.Println(err)
	}
}

// is that needed?

func Create(tx *gorm.DB, o any) error {
	return db.Create(tx, o)
}
func Read(tx *gorm.DB, o any) error {
	return db.Read(tx, o)
}
func Update(tx *gorm.DB, o any) error {
	return db.Update(tx, o)
}
func Delete(tx *gorm.DB, o any) error {
	return db.Delete(tx, o)
}

const (
	MyRelationStatusRequested = "requested"
	MyRelationStatusAccepted  = "accepted"
	MyRelationStatusBlocking  = "blocking"
)

const (
	StatusVisibilityPublic   = "public"
	StatusVisibilityUnlisted = "unlisted"
	StatusVisibilityPrivate  = "private"
	StatusVisibilityDirect   = "direct"
)

type MyAccount struct {
	// meta
	Email      string `json:"email" gorm:"index:email;type:text collate nocase"`
	PasswdHash string `json:"-"`

	// without @host
	Username string `json:"preferredUsername" gorm:"primarykey;index:username;type:text collate nocase"`
	Host     string `gorm:"primarykey;index:host;type:text collate nocase"`

	// activitypub url
	ActivitypubID string `gorm:"index:activitypub_id,unique;type:text collate nocase"`
	// timestamp string
	AccountID string `gorm:"index:mastodon_id,unique;type:text collate nocase"`

	AlsoKnownAs []string `json:"alsoKnownAs,omitempty" gorm:"serializer:json"`

	PrivateKeyPem string

	CreatedAt int64
	DeletedAt int64
}

// actor follows object.
// all id is activitypub id.
type MyRelation struct {
	ID int `gorm:"primarykey;autoIncrement"`

	Actor  string `gorm:"uniqueIndex:idx_actor_object"`
	Object string `gorm:"uniqueIndex:idx_actor_object"`

	// "requested" | "accepted" | "blocking"
	Status string
}

// all id is activitypub id.
type MyReblog struct {
	ID         string `gorm:"primarykey"`
	Actor      string
	Object     string
	Visibility string

	// the owner of the object, for receive the notification
	Owner string

	DeletedAt int64
}

// all id is activitypub id.
type MyFavourite struct {
	ID     int    `gorm:"primarykey;autoIncrement"`
	Actor  string `gorm:"uniqueIndex:idx_actor_object"`
	Object string `gorm:"uniqueIndex:idx_actor_object"`

	// the owner of the object, for receive the notification
	Owner string

	DeletedAt int64
}

// utils

func logE(v ...any) {
	log.Println(v...)
}
