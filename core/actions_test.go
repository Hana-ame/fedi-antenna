package core

import (
	"testing"
)

func TestFollow(t *testing.T) {
	err := Follow("nanaka", "fedi.moonchan.xyz", "https://mstdn.social/users/tsukishima_test")
	if err != nil {
		t.Error(err)
	}
}
