package model

import (
	"fmt"

	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/orderedmap"
)

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
		utils.NewMapFromKV([]*utils.KV{
			{Key: "rel", Value: "self"},
			{Key: "type", Value: "application/activity+json"},
			{Key: "href", Value: "https://" + host + "/users/" + username},
		}),
		utils.NewMapFromKV([]*utils.KV{
			{Key: "rel", Value: "http://webfinger.net/rel/profile-page"},
			{Key: "type", Value: "text/html"},
			{Key: "href", Value: "https://" + host + "/@" + username},
		}),
		utils.NewMapFromKV([]*utils.KV{ // dunno what it is
			{Key: "rel", Value: "http://ostatus.org/schema/1.0/subscribe"},
			{Key: "template", Value: "https://p1.a9z.dev/authorize-follow?acct={uri}"},
		}),
	})
	return
}
