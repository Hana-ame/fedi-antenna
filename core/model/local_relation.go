package model

// following / follower
// padding
// blocking

type LocalRelation struct {
	// object ID in url
	ID string `gorm:"primarykey"`
	// activitypub ID in url
	Actor  string
	Object string
	// Follow or Block
	Type string
	// padding or acccpted
	Status string
}

const (
	RelationTypeFollow = "Follow"
	RelationTypeBlock  = "Block"
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
