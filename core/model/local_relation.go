package model

// following / follower
// padding
// blocking

type LocalRelation struct {
	// object ID in url
	ID string
	// activitypub ID in url
	Actor  string `gorm:"primarykey"`
	Object string `gorm:"primarykey"`
	// Follow or Block
	Type string
	// padding or acccpted
	Status string
}

const (
	RelationTypeFollow = "Follow"
	RelationTypeBlock  = "Block"
	RelationTypeNone   = ""
)

const (
	RelationStatusPadding  = "padding"
	RelationStatusAccepted = "accepted"
	RelationStatusRejected = "rejected"

	RelationStatusUnblocked = "unblocked"
	RelationStatusBlocked   = "blocked"
	RelationStatusBlocking  = "blocking"

	RelationStatusUndo = "undo"
)
