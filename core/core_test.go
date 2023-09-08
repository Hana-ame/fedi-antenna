package core

import (
	"fmt"
	"testing"
)

func TestGetRemoteUser(t *testing.T) {
	user, err := GetRemoteUser("misRoute@mona.do")
	fmt.Println(user)
	fmt.Println(err)
}

func TestGetRemoteUser2(t *testing.T) {
	user, err := GetRemoteUser("misRoute@mona.do")
	fmt.Println(user)
	fmt.Println(err)
}
