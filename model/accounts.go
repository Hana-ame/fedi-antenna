package model

type Account struct {
	// activitypub
	ID string `gorm:"primarykey"`

	Timestamp int64
	// user@host
	Account string
}
