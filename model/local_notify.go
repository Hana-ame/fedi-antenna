package model

// reblogs are also here
// in actor = actor, type = Announce
// find in noteID
type LocalNotify struct {
	// statuses's ID in url
	// favourite：objectID
	// reblog：statusID + /acitivy
	ID string `gorm:"primarykey"`
	// activitypub ID in url, actor is sender
	Actor string
	// activitypub ID in url
	// the note that reblogged/favourited
	Object string
	// Reblog or Favourite or Mention?
	Type string

	// activitypub ID in url, for the receiver.
	// the owner of the object
	To string

	// reblog only, public, unlisted, private
	Visibility string

	// todo
	DeleteAt int64
}

const (
	NotifyTypeAnnounce = "Announce"
	NotifyTypeLike     = "Like"
	NotifyTypeMention  = "Mention" //?
)
