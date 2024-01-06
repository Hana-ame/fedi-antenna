package main

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/Hana-ame/fedi-antenna/activitypub/actions"
	"github.com/Hana-ame/fedi-antenna/activitypub/model"
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

// pass 231228
func TestFollow(t *testing.T) {
	actor := "https://fedi.moonchan.xyz/users/test5"
	object := "https://me.ns.ci/users/cocoon"
	err := actions.Follow(actor, object)
	log.Println(err)
	_ = err
}

func TestUndo(t *testing.T) {
	actor := "https://fedi.moonchan.xyz/users/test1"
	object := "https://fedi.moonchan.xyz/objects/follow/0f8de8f8-c199-45fc-bbcd-c08b843df7fe"
	err := actions.UndoFollow(actor, object)
	log.Println(err)
}

func TestAccept(t *testing.T) {
	actor := "https://fedi.moonchan.xyz/users/test5"
	object := "https://me.ns.ci/users/cocoon"
	fmt.Println(object)
	err := actions.Accept(actor, object)
	log.Println(err)
}
func TestReject(t *testing.T) {
	actor := "https://fedi.moonchan.xyz/users/test5"
	object := "https://me.ns.ci/users/cocoon"
	fmt.Println(object)
	err := actions.Reject(actor, object)
	log.Println(err)
}

func TestNote(t *testing.T) {
	actor := utils.ParseActivitypubID("test5", "fedi.moonchan.xyz")

	timestamp := utils.Now()
	published := utils.TimestampToRFC3339(timestamp)
	name, host := utils.ParseNameAndHost(actor)
	note := &activitypub.Note{
		ID:          utils.ParseStatusesID(name, host, strconv.Itoa(int(timestamp))),
		Summary:     utils.ParseStringToPointer("", true),
		Content:     "你们能不能去死啊，吵死了",
		Visibility:  model.VisiblityPublic,
		InReplyTo:   utils.ParseStringToPointer("", true),
		AttributeTo: actor,
		Published:   published,
	}
	note.Autofill()
	err := actions.CreateNote(note)

	log.Printf("%s", err)
}
