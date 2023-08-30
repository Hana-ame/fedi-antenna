package core

import (
	"fmt"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/activitypub"
	"github.com/Hana-ame/fedi-antenna/core/webfinger"
	"github.com/iancoleman/orderedmap"
)

func Inbox(header http.Header, o *orderedmap.OrderedMap, verify error) {
	fmt.Println(verify)
}

func IsUserExist(username, host string) bool {
	if username == "nanaka" || username == "misRoute" {
		return true
	}
	return false
}

// user
func APUserObj(user, host string) (o *orderedmap.OrderedMap) {
	var published int64 = 1693394962808
	pubkey := activitypub.PublicKeyObj(user, host)
	icon := activitypub.ImageObj("image/jpeg", "https://s3.arkjp.net/misskey/678ad158-f160-48f4-a369-8756aa92350e.jpg")
	o = activitypub.UserObj(
		host, user,
		published, // timestamp in us,
		pubkey, icon,
	)
	return
}

// webfinger
func CreateWebfingerObj(username, host string) (o *orderedmap.OrderedMap, err error) {
	return webfinger.CreateWebfingerObj(username, host)
}
