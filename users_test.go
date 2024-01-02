package main

import (
	"fmt"
	"testing"

	"github.com/Hana-ame/fedi-antenna/core"
)

func TestUser(t *testing.T) {
	id := "https://fedi.moonchan.xyz/users/test1"
	user, err := core.ReadActivitypubUserByID(id)
	fmt.Println(user)
	fmt.Println(err)
}
