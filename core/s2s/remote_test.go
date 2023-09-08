package s2s

import (
	"fmt"
	"testing"
)

// pass
func TestGetRemoteUser(t *testing.T) {
	// u, err := GetRemoteUser("misRoute@mona.do")
	u, err := GetRemoteUser("meromero@p1.a9z.dev")
	fmt.Println(u, err)
}
