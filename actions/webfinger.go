package actions

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/Tools/myfetch"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/webfinger/model"

	"github.com/Hana-ame/orderedmap"
)

func FetchWebfingerByAcct(acct string) (o *orderedmap.OrderedMap, err error) {
	accountID := &model.AccountID{
		Account: acct,
	}
	if err := dao.Read(accountID); err == nil && accountID.WebfingerObject != nil {
		return accountID.WebfingerObject, err
	}
	username, host := utils.ParseUserAndHost(acct)
	url := utils.ParseWebfingerUrl(username, host)
	resp, err := myfetch.Fetch(http.MethodGet, url, nil, nil)
	if err != nil {
		return
	}
	o, err = myfetch.ResponseToObject(resp)
	accountID.WebfingerObject = o
	dao.Update(accountID)

	return
}
