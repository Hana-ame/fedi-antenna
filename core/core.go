package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/activitypub"
	"github.com/Hana-ame/fedi-antenna/core/s2s"
	"github.com/Hana-ame/fedi-antenna/core/webfinger"
	"github.com/Hana-ame/fedi-antenna/db"
	"github.com/Hana-ame/fedi-antenna/utils"
	"github.com/iancoleman/orderedmap"
)

// not done
func Inbox(header http.Header, o *orderedmap.OrderedMap, verify error) {
	headerBytes, _ := json.Marshal(header)
	headerStr := string(headerBytes)
	bodyBytes, _ := json.Marshal(o)
	bodyStr := string(bodyBytes)
	verifyStr := fmt.Sprintf("%s", verify)
	err := db.CreateLog(&headerStr, &bodyStr, &verifyStr)
	if err != nil {
		fmt.Printf("%s", err)
	}
}

// test
// todo
// user
// func APUserObj(user, host string) (o *orderedmap.OrderedMap) {
// 	var published int64 = 1693394962808
// 	pubkey := activitypub.PublicKeyObj(user, host, )
// 	icon := activitypub.ImageObj("image/jpeg", "https://s3.arkjp.net/misskey/678ad158-f160-48f4-a369-8756aa92350e.jpg")
// 	o = activitypub.UserObj(
// 		host, user,
// 		published, // timestamp in us,
// 		pubkey, icon,
// 	)
// 	return
// }

// webfinger

// not used
// func IsUserExist(username, host string) bool {
// 	_, err := db.ReadUserByKey(username, host)
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }

func WebfingerObj(username, host string) (o *orderedmap.OrderedMap, err error) {
	user, err := db.ReadUserByKey(username, host)
	if err != nil {
		return nil, err
	}
	return webfinger.CreateWebfingerObj(user.Username, user.Host)
}

func UserObj(username, host string) (o *orderedmap.OrderedMap, err error) {
	user, err := db.ReadUserByKey(username, host)
	if err != nil {
		return nil, err
	}

	pk, err := utils.ParsePrivateKey([]byte(user.PrivateKeyPem))
	if err != nil {
		return nil, err
	}

	pubkeyPem := utils.MarshalPublicKey(&pk.PublicKey)
	pubkey := activitypub.PublicKeyObj(user.Username, user.Host, string(pubkeyPem))
	icon := activitypub.ImageObj("image/jpeg", "https://s3.arkjp.net/misskey/678ad158-f160-48f4-a369-8756aa92350e.jpg") // TEMP

	o = activitypub.UserObj(
		user.Host, user.Username, user.Published,
		pubkey, icon,
	)
	return o, nil
}

// s2s

// get remote user and save it into db
func CacheRemoteUser(acct string) error {
	user, err := s2s.GetRemoteUser(acct)
	if err != nil {
		return err
	}

	// update user or create user
	if err := db.UpdateRemoteUser(user); err != nil {
		err = db.CreateRemoteUser(user)
		if err != nil {
			return err
		}
	}

	return nil
}

// get remote user from db, if not exist, run CacheRemoteUser
func GetRemoteUser(acct string) (*db.RemoteUser, error) {
	user := &db.RemoteUser{Acct: acct}
	if err := db.ReadRemoteUser(user); err == nil {
		if user.LastSeen+1000*60*60*6 < utils.TimestampInMs() { // 6 hours
			go CacheRemoteUser(acct)
		}
		return user, nil // success
	}
	if err := CacheRemoteUser(acct); err != nil {
		return user, err // err
	}
	if err := db.ReadRemoteUser(user); err != nil {
		return user, err // err
	}
	return user, nil // success
}

// local user

// create local user
func CreateUser(user *db.User) error {
	if user.Attatchment == "" {
		user.Attatchment = "[]"
	}
	if user.AlsoKnownAs == "" {
		user.AlsoKnownAs = "[]"
	}
	user.Published = utils.TimestampInMs()
	user.PrivateKeyPem = string(utils.MarshalPrivateKey(utils.GenerateKey()))

	if err := db.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func DeleteUser(user *db.User) error {
	if err := db.ReadUser(user); err != nil {
		return err
	}

	user.Deleted = true

	if err := db.UpdateUser(user); err != nil {
		return err
	}

	return nil
}

// local action
type ApiAction struct {
	Action string `json:"action"`
	Actor  string `json:"actor"`
	Object string `json:"object"`
}
