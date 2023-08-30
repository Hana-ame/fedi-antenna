package webfinger

import (
	"fmt"
	"testing"
)

func TestFetchWebfingerObj(t *testing.T) {
	o, err := FetchWebfingerObj("meromero@p1.a9z.dev")
	fmt.Println(o, err)
}

func TestGetUserIdFromAcct(t *testing.T) {
	a, err := GetUserIdFromAcct("meromero@p1.a9z.dev")
	fmt.Println(a, err)
}
