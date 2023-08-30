package utils

import (
	"errors"
	"strings"
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
