package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/Hana-ame/fedi-antenna/core/actions"
)

// pass 231228
func TestFollow(t *testing.T) {
	actor := "https://fedi.moonchan.xyz/users/test1"
	object := "https://p1.a9z.dev/users/9a3qtdtypj"
	// object := "https://p1.a9z.dev/users/9a3qtdtypj"
	err := actions.Follow(actor, object)
	// log.Println(err)
	_ = err
}

func TestUndo(t *testing.T) {
	actor := "https://fedi.moonchan.xyz/users/test1"
	object := "https://fedi.moonchan.xyz/objects/follow/a15d7b20-c5c0-4f88-94a7-c42c75a9c7fa"
	err := actions.UndoFollow(actor, object)
	log.Println(err)
}

func TestAccept(t *testing.T) {
	actor := "https://fedi.moonchan.xyz/users/test1"
	object := "https://p1.a9z.dev/follows/9nt3kg7k4d"
	fmt.Println(object)
	err := actions.Accept(actor, object)
	log.Println(err)
}
