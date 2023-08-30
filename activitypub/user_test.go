package activitypub

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestU(t *testing.T) {
	user := "nanaka"
	host := "fedi.moonchan.xyz"
	var published int64 = 1693394962808
	pubkey := PublicKeyObj(user, host)
	icon := ImageObj("image/jpeg", "https://s3.arkjp.net/misskey/678ad158-f160-48f4-a369-8756aa92350e.jpg")
	o := UserObj(
		host, user,
		published, // timestamp in us,
		pubkey, icon,
	)
	fmt.Println(o)
	b, _ := json.Marshal(o)
	fmt.Printf("%s", b)
}
