package actions

import (
	"fmt"

	"github.com/Hana-ame/fedi-antenna/core/utils"
)

// acct = "meromero@p1.a9z.dev"
func GetUserIdFromAcct(acct string, skipCache bool) (userId string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%s", e)
		}
	}()

	o, err := FetchWebfingerByAcct(acct, skipCache)
	userId = utils.ParseUserId(o)

	return
}
