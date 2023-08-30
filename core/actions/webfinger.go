package actions

import (
	"fmt"

	"github.com/Hana-ame/fedi-antenna/core/webfinger"
)

// acct = "meromero@p1.a9z.dev"
func GetUserIdFromAcct(acct string) (userId string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%s", e)
		}
	}()
	o, err := webfinger.FetchWebfingerObj(acct)
	handleErr(err)

	userId = webfinger.ParseUserId(o)

	return
}
