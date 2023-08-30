package webfinger

import (
	"fmt"
	"testing"
)

func TestFetchWebfingerObj(t *testing.T) {
	o, err := FetchWebfingerObj("meromero@p1.a9z.dev")
	fmt.Println(o, err)
}
