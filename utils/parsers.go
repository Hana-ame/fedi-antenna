package utils

import (
	"log"
	"strings"
)

// parse

// `a@b.com` => "a", "b.com"
// or
// `@a@b.com` => "a", "b.com"
func ParseUserAndHost(acct string) (user, host string) {
	acct = strings.TrimPrefix(acct, "@")
	arr := strings.Split(acct, "@")
	if len(arr) != 2 {
		log.Println("the format of acct is incorrect:" + acct)
		return
	}
	user, host = arr[0], arr[1]
	return
}

// func ParseObjValueToString(o *orderedmap.OrderedMap, key string) (string, bool) {
// 	if v, ok := o.Get(key); !ok {
// 		return "not found", false
// 	} else {
// 		s, ok := v.(string)
// 		return s, ok
// 	}
// }

// just add @ between [user] and [host]
func ParseAcctStr(user, host string) (acct string) {
	return user + "@" + host
}

// "meromero", "p1.a9z.dev"	=> "https://p1.a9z.dev/.well-known/webfinger?resource=acct:meromero@p1.a9z.dev"
func ParseWebfingerUrl(username, host string) string {
	return `https://` + host + `/.well-known/webfinger?resource=acct:` + username + `@` + host
}
