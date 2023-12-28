package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/Hana-ame/fedi-antenna/core/actions"
)

// pass 231228
func TestFollow(t *testing.T) {
	fmt.Println("donotca1che")
	actor := "https://fedi.moonchan.xyz/users/test1"
	object := "https://p1.a9z.dev/users/9a3qtdtypj"
	err := actions.Follow(actor, object)
	// log.Println(err)
	_ = err
}

func TestUndo(t *testing.T) {
	actor := "https://fedi.moonchan.xyz/users/test1"
	object := "https://fedi.moonchan.xyz/objects/follow/0f8de8f8-c199-45fc-bbcd-c08b843df7fe"
	err := actions.UndoFollow(actor, object)
	log.Println(err)
}

func TestAccept(t *testing.T) {
	actor := "https://fedi.moonchan.xyz/users/test1"
	object := "https://p1.a9z.dev/follows/9nt4sxsk4s"
	fmt.Println(object)
	err := actions.Accept(actor, object)
	log.Println(err)
}
func TestReject(t *testing.T) {
	actor := "https://fedi.moonchan.xyz/users/test1"
	object := "https://p1.a9z.dev/follows/9nt4sxsk4s"
	fmt.Println(object)
	err := actions.Reject(actor, object)
	log.Println(err)
}
