package main

import (
	"log"
	"testing"

	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
)

// passed
func TestNewUser(t *testing.T) {
	user := activitypub.NewUser("test1", "fedi.moonchan.xyz")
	log.Println(user)
	err := dao.Create(user)
	log.Println(err)
}
