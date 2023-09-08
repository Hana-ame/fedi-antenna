package s2s

import (
	"encoding/json"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/core/webfinger"
	"github.com/Hana-ame/fedi-antenna/db"
	"github.com/Hana-ame/fedi-antenna/utils"
)

func GetRemoteUser(acct string) (*db.RemoteUser, error) {
	id, err := webfinger.GetUserIdFromAcct(acct)
	if err != nil {
		return nil, err
	}

	o, err := utils.FetchObj(http.MethodGet, id, nil, nil)
	if err != nil {
		return nil, err
	}

	oBytes, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	user := &db.RemoteUser{ID: id, Acct: acct, O: string(oBytes), LastSeen: utils.TimestampInMs()}

	return user, nil
}
