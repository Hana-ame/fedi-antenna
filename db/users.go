package db

// local users

type User struct {
	Username                  string `gorm:"primaryKey"` // [username]@host.com
	Host                      string `gorm:"primaryKey"` // username@[host.com]
	Name                      string // that can be set
	Summary                   string // 简介
	ManuallyApprovesFollowers bool   //
	Published                 int64  // in ms, created at
	AlsoKnownAs               string // []string
	Attatchment               string // []Attatchment
	Icon                      string // activitypub.ImageObj marshaled

	Deleted bool // deleted
}

// ap users

// 大体上是个缓存。、
type APUser struct {
	ID       string `gorm:"primaryKey"`        // id: https://host.com/users/user
	Acct     string `gorm:"index:acct,unique"` // user@host.com
	O        string // user object
	LastSeen int64  // timestamp in ms
}

// methods

// local users

// ap users
