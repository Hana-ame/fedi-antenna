package webfinger

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/utils"
	"github.com/iancoleman/orderedmap"
)

// acct = "meromero@p1.a9z.dev"
func GetUserIdFromAcct(acct string) (userId string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%s", e)
		}
	}()
	o, err := FetchWebfingerObj(acct)
	handleErr(err)

	userId = ParseUserId(o)

	return
}

// as server

func CreateWebfingerObj(username, host string) (o *orderedmap.OrderedMap, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%s", e)
		}
	}()
	o = orderedmap.New()
	o.Set("subject", "acct:"+username+"@"+host)
	o.Set("links", []*orderedmap.OrderedMap{
		utils.CreateOrderedMap([]*utils.KV{
			{Key: "rel", Value: "self"},
			{Key: "type", Value: "application/activity+json"},
			{Key: "href", Value: "https://" + host + "/users/" + username},
		}),
		utils.CreateOrderedMap([]*utils.KV{
			{Key: "rel", Value: "http://webfinger.net/rel/profile-page"},
			{Key: "type", Value: "text/html"},
			{Key: "href", Value: "https://" + host + "/@" + username},
		}),
		utils.CreateOrderedMap([]*utils.KV{ // dunno what it is
			{Key: "rel", Value: "http://ostatus.org/schema/1.0/subscribe"},
			{Key: "template", Value: "https://p1.a9z.dev/authorize-follow?acct={uri}"},
		}),
	})
	return
}

// as client

func FetchWebfingerObj(acct string) (o *orderedmap.OrderedMap, err error) {
	username, host := utils.ParseAcctStrToUserAndHost(acct)
	url := ParseWebfingerUrl(username, host)
	o, err = utils.FetchObj(http.MethodGet, url, nil)
	return
}

// "meromero", "p1.a9z.dev"	=> "https://p1.a9z.dev/.well-known/webfinger?resource=acct:meromero@p1.a9z.dev"
func ParseWebfingerUrl(username, host string) string {
	return `https://` + host + `/.well-known/webfinger?resource=acct:` + username + `@` + host
}

// Parse UserId from webfinger
func ParseUserId(webfingerObj *orderedmap.OrderedMap) string {
	if links, ok := webfingerObj.Get("links"); !ok {
		handleErr(errors.New("notfound links"))
	} else {
		if arr, ok := links.([]any); !ok {
			handleErr(errors.New("links not array"))
		} else {
			for _, li := range arr {
				if lo, ok := li.(orderedmap.OrderedMap); ok {
					key, ok := lo.Get("rel")
					if ok && key == "self" {
						id, ok := lo.Get("href")
						if !ok {
							handleErr(errors.New("self link doesn't have href"))
						} else {
							return id.(string)
						}
					}
				}
			}
		}
	}
	handleErr(errors.New("id not found"))
	return "" // never
}

// utils
func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
