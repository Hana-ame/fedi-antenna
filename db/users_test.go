package db

import (
	"fmt"
	"testing"

	"github.com/Hana-ame/fedi-antenna/utils"
)

func TestRemoteUser(t *testing.T) {
	Init()
	user := &RemoteUser{
		ID:       "2",
		Acct:     "2",
		O:        utils.TimestampToRFC3339(utils.TimestampInMs()),
		LastSeen: utils.TimestampInMs(),
	}
	err := CreateRemoteUser(user)
	fmt.Println(err)
}

func TestUpdateRemoteUser(t *testing.T) {
	Init()
	user := &RemoteUser{
		ID:       "1",
		Acct:     "1",
		O:        utils.TimestampToRFC3339(utils.TimestampInMs()),
		LastSeen: utils.TimestampInMs(),
	}
	err := UpdateRemoteUser(user)
	fmt.Println(err)
}
