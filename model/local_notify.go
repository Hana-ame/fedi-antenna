package model

type LocalNotify struct {
	// statuses's ID in url
	ID string `gorm:"primarykey"`
	// activitypub ID in url, actor is sender
	Actor string `gorm:"primarykey"`
	// activitypub ID in url, object is reciver
	Object string
	// Reblog or Favourite or Mention
	Type string `gorm:"primarykey"`
}

const (
	NotifyTypeReblog    = "Reblog"
	NotifyTypeFavourite = "Favourite"
	NotifyTypeMention   = "Mention"
)
