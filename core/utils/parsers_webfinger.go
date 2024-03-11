package utils

import (
	"log"
	"strings"

	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
)

// parse

// `a@b.com` => "a", "b.com"
// or
// `@a@b.com` => "a", "b.com"
func ParseUserAndHost(acct string) (user, host string) {
	acct = strings.TrimPrefix(acct, "@")
	arr := strings.Split(acct, "@")
	if len(arr) < 2 {
		log.Println("the format of acct is incorrect:" + acct)
		return
	}
	user, host = arr[0], arr[1]
	return
}

// just add @ between [user] and [host]
func UsernameAndHost2Account(user, host string) (acct string) {
	return user + "@" + host
}

// "meromero", "p1.a9z.dev"	=> "https://p1.a9z.dev/.well-known/webfinger?resource=acct:meromero@p1.a9z.dev"
func ParseWebfingerUrl(username, host string) string {
	return `https://` + host + `/.well-known/webfinger?resource=acct:` + username + `@` + host
}

// too old
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
