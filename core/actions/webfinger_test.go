package actions

import (
	"fmt"
	"testing"
)

func TestGetUserIdFromAcct(t *testing.T) {
	a, err := GetUserIdFromAcct("meromero@p1.a9z.dev")
	fmt.Println(a, err)
}
