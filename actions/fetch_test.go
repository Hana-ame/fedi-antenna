package actions

import (
	"fmt"
	"io"
	"testing"
)

func TestFetch(t *testing.T) {
	r, _ := Fetch("GET", "https://p1.a9z.dev/users/9a3qtdtypj", nil, nil)
	h, _ := io.ReadAll(r.Body)
	fmt.Println(h)

}
