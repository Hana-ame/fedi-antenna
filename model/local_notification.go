package model

// reblogs are also here
// in actor = actor, type = Announce
// find in noteID
type LocalNotify struct {
	// statuses's ID in url
	ID string `gorm:"primarykey"`
	// activitypub ID in url, actor is sender
	Actor string
	// activitypub ID in url
	Object string
	// Reblog or Favourite or Mention?
	Type string

	// activitypub ID in url, for the receiver.
	To string

	Visibility string

	DeleteAt int64
}

const (
	NotifyTypeAnnounce = "Announce"
	NotifyTypeLike     = "Like"
	NotifyTypeMention  = "Mention" //?
)
