package main

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/Hana-ame/fedi-antenna/actions"
	"github.com/Hana-ame/fedi-antenna/actions/model"
	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
	core "github.com/Hana-ame/fedi-antenna/core/model"
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
	actor := "https://fedi.moonchan.xyz/users/test7"
	object := "https://me.ns.ci/users/cocoon"
	fmt.Println(object)
	lr := &core.LocalRelation{
		Actor:  object,
		Object: actor,
	}
	err := actions.Accept(lr, true)
	log.Println(err)
}
func TestReject(t *testing.T) {
	actor := "https://fedi.moonchan.xyz/users/test5"
	object := "https://me.ns.ci/users/cocoon"
	fmt.Println(object)
	// lr := &core.LocalRelation{
	// 	Actor:  object,
	// 	Object: actor,
	// }
	err := actions.Reject(object, actor)
	log.Println(err)
}

func TestNote(t *testing.T) {
	actor := utils.NameAndHost2ProfileUrlActivitypubID("test7", "fedi.moonchan.xyz")

	timestamp := utils.NewTimestamp()
	published := utils.TimestampToRFC3339(timestamp)
	name, host := utils.ActivitypubID2NameAndHost(actor)
	note := &activitypub.Note{
		ID:           utils.ParseStatusesUri(name, host, strconv.Itoa(int(timestamp))),
		Summary:      utils.ParseStringToPointer("", true),
		Content:      "要先follow",
		Visibility:   model.VisiblityPublic,
		InReplyTo:    utils.ParseStringToPointer("", true),
		AttributedTo: actor,
		Published:    published,
	}
	note.Autofill()
	err := actions.CreateNote(note)

	log.Printf("%s", err)
}
