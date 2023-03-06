package webfinger

import (
	"fmt"
	"strings"
)

// just a helper function for finding self.href from resource object
func GetIdFromResource(resource *Resource) (string, error) {
	for _, link := range resource.Links {
		if link.Rel == "self" {
			return link.HRef, nil
		}
	}
	return "", fmt.Errorf("not webfinger, find no links with rel=self")
}

func ParseAcct(acct string) (username string, domain string) {
	// if strings.HasPrefix(acct, "acct:") {
	// 	acct = acct[5:]
	// }
	acct = strings.TrimPrefix(acct, "acct:")

	arr := strings.Split(acct, "@")
	if len(arr) == 2 {
		return arr[0], arr[1]
	}
	return acct, ""
}

func CheckUserExist(username string) bool {
	return true
}
