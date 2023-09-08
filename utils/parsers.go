package utils

import (
	"errors"
	"strings"

	"github.com/iancoleman/orderedmap"
)

// parse

// "a@b.com" => "a", "b.com"
//
// "@a@b.com" => "a", "b.com"
func ParseAcctStrToUserAndHost(acct string) (user, host string) {
	acct = strings.TrimPrefix(acct, "@")
	arr := strings.Split(acct, "@")
	if len(arr) != 2 {
		handleErr(errors.New("the format of acct is incorrect:" + acct))
	}
	user, host = arr[0], arr[1]
	return
}

func ParseObjValueToString(o *orderedmap.OrderedMap, key string) (string, bool) {
	if v, ok := o.Get(key); !ok {
		return "not found", false
	} else {
		s, ok := v.(string)
		return s, ok
	}
}

func ParseNameAndHostToAcctStr(user, host string) (acct string) {
	return user + "@" + host
}
