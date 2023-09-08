package db

// local users

type User struct {
	Username                  string `gorm:"primaryKey;type:text collate nocase"` // [username]@host.com
	Host                      string `gorm:"primaryKey;type:text collate nocase"` // username@[host.com]
	Name                      string // that can be set
	Summary                   string // 简介
	ManuallyApprovesFollowers bool   //
	Published                 int64  // in ms, created at
	AlsoKnownAs               string // []string
	Attatchment               string // []Attatchment
	Icon                      string // activitypub.ImageObj marshaled

	Deleted bool // deleted

	Passhash      string // the hash of User's passwd
	PrivateKeyPem string // privatekey
}

// remote users

// 大体上是个缓存。、
type RemoteUser struct {
	ID       string `gorm:"primaryKey;type:text collate nocase"`        // id: https://host.com/users/user
	Acct     string `gorm:"index:acct,unique;type:text collate nocase"` // user@host.com
	O        string // user object
	LastSeen int64  // timestamp in ms
}

// methods

// local users

func CreateUser(user *User) error {
	tx := db.Create(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func ReadUserByKey(username, host string) (*User, error) {
	user := &User{Username: username, Host: host}
	tx := db.Take(user)
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func ReadUser(user *User) error {
	// user := &User{Username: username, Host: host}
	tx := db.Take(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateUser(user *User) error {
	tx := db.Save(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// ap users
func CreateRemoteUser(user *RemoteUser) error {
	tx := db.Create(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func ReadRemoteUser(user *RemoteUser) error {
	// user := &RemoteUser{ID: id}
	// user := &RemoteUser{Acct: acct}
	tx := db.Take(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateRemoteUser(user *RemoteUser) error {
	tx := db.Save(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
