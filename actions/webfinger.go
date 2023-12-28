package actions

import (
	"fmt"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/Tools/myfetch"
	"github.com/Hana-ame/fedi-antenna/core/utils"

	"github.com/Hana-ame/orderedmap"
)

func FetchWebfingerByAcct(acct string) (o *orderedmap.OrderedMap, err error) {
	username, host := utils.ParseUserAndHost(acct)
	url := utils.ParseWebfingerUrl(username, host)
	resp, err := myfetch.Fetch(http.MethodGet, url, nil, nil)
	if err != nil {
		return
	}
	o, err = myfetch.ResponseToObject(resp)

	return
}

// acct = "meromero@p1.a9z.dev"
func GetUserIdFromAcct(acct string) (userId string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%s", e)
		}
	}()
	o, err := FetchWebfingerByAcct(acct)
	userId = utils.ParseUserId(o)

	return
}
