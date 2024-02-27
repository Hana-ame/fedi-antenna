package model

// this used to recordf what fediverse
type FediStatus struct {
	Host        string `json:"host" gorm:"primarykey"`
	SharedInbox string

	LastSeen int64 `json:"last_seen"`
}
