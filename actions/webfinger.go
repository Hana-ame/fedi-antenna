package actions

import (
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
