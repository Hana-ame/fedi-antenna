package curl

import (
	"fmt"
	"testing"
)

func TestWebfinger(t *testing.T) {
	o, e := Webfinger("nanakananoka@mstdn.jp")
	fmt.Println(e)
	fmt.Println(o)
}
