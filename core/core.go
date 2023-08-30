package core

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/core/webfinger"
	"github.com/iancoleman/orderedmap"
)

func Inbox(header http.Header, o *orderedmap.OrderedMap, verify error) {

}

func IsUserExist(username, host string) bool {
	if username == "nanaka" || username == "misRoute" {
		return true
	}
	return false
}

// webfinger
func CreateWebfingerObj(username, host string) (o *orderedmap.OrderedMap, err error) {
	return webfinger.CreateWebfingerObj(username, host)
}
