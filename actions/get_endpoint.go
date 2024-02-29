package actions

import (
	"fmt"
	"strings"

	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

const activitystreamsPublic = "https://www.w3.org/ns/activitystreams#Public"

// 草，要么得改表，这段代码有大病。
// 为啥这么写，因为我懒。

// return inboxex
func Endpoints(to, cc []string) (endpoints []string) {
	for _, v := range to {
		if v == activitystreamsPublic {
			return nil
		}
	}
	for _, v := range cc {
		if v == activitystreamsPublic {
			return nil
		}
	}
	endpoints = make([]string, 0)
	for _, v := range to {
		if strings.HasSuffix(v, "followers") {
			s, _ := findEndpoints(v)
			if s != nil {
				for _, ss := range s {
					endpoints = append(endpoints, ss)
				}
			}
		}
	}
	for _, v := range cc {
		if strings.HasSuffix(v, "followers") {
			s, _ := findEndpoints(v)
			if s != nil {
				for _, ss := range s {
					endpoints = append(endpoints, ss)
				}
			}
		}
	}
	return endpoints
}

// input id may contain follow.
// not testd
func findEndpoints(id string) (endpoints []string, err error) {
	if !(strings.HasSuffix(id, "follower") || strings.HasSuffix(id, "followers")) {
		if endpoint, err := FindEndpointsByID(id); err != nil {
			return nil, err
		} else {
			return []string{endpoint}, fmt.Errorf("not follower endpint")
		}
	}
	// local folloers entry point
	name, host := utils.ActivitypubID2NameAndHost(id)

	endpoints, err = dao.ReadFollowersByLocaluserID(utils.NameAndHost2ProfileUrlActivitypubID(name, host))

	return
}

// find from local,?
func FindEndpointsByID(id string) (endpoint string, err error) {
	// user := &activitypub.User{
	// 	ID: id,
	// }
	user, err := core.ReadActivitypubUserByID(id)
	if err != nil {
		return "", err
	}
	endpoint = user.Inbox
	return

}
