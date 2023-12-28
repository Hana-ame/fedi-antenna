package actions

import (
	"fmt"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/Tools/myfetch"
	"github.com/Hana-ame/fedi-antenna/core/utils"

	"github.com/Hana-ame/orderedmap"
)

func FetchWebfingerByAcct(acct string) (o *orderedmap.OrderedMap) {
	return
}

// as client

func FetchWebfingerObj(acct string) (o *orderedmap.OrderedMap, err error) {
	username, host := utils.ParseUserAndHost(acct)
	url := utils.ParseWebfingerUrl(username, host)
	resp, err := myfetch.Fetch(http.MethodGet, url, nil, nil)
	if err != nil {
		return
	}
	o, err = myfetch.ResponseToObject(resp)

	return
}

// Parse UserId from webfinger
func ParseUserId(webfingerObj *orderedmap.OrderedMap) string {
	if links, ok := webfingerObj.Get("links"); !ok {
		return ""
	} else {
		if arr, ok := links.([]any); !ok {
			return ""
		} else {
			for _, li := range arr {
				if lo, ok := li.(orderedmap.OrderedMap); ok {
					key, ok := lo.Get("rel")
					if ok && key == "self" {
						id, ok := lo.Get("href")
						if !ok {
							return ""
						} else {
							return id.(string)
						}
					}
				}
			}
		}
	}
	return "" // never
}

// acct = "meromero@p1.a9z.dev"
func GetUserIdFromAcct(acct string) (userId string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%s", e)
		}
	}()
	o, err := FetchWebfingerObj(acct)
	userId = ParseUserId(o)

	return
}
