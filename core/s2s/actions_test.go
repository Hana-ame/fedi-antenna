package s2s

import (
	"encoding/json"
	"testing"

	"github.com/iancoleman/orderedmap"
)

func TestFollow(t *testing.T) {
	err := Follow("nanaka", "fedi.moonchan.xyz", "https://mstdn.social/users/tsukishima_test")
	if err != nil {
		t.Error(err)
	}
}
func TestReject(t *testing.T) {
	data := `{"@context":"https://www.w3.org/ns/activitystreams","id":"https://mstdn.social/4ee8c5aa-88d7-4328-a9d6-314a1f769956","type":"Follow","actor":"https://mstdn.social/users/tsukishima_test","object":"https://fedi.moonchan.xyz/users/nanaka"}`
	o := orderedmap.New()
	json.Unmarshal([]byte(data), o)
	err := Reject("nanaka", "fedi.moonchan.xyz", o)
	if err != nil {
		t.Error(err)
	}
}
func TestAccept(t *testing.T) {
	data := `{"@context":"https://www.w3.org/ns/activitystreams","id":"https://mstdn.social/48abc96b-5485-4e5b-9e62-26d9413e3d95","type":"Follow","actor":"https://mstdn.social/users/tsukishima_test","object":"https://fedi.moonchan.xyz/users/nanaka"}`
	o := orderedmap.New()
	json.Unmarshal([]byte(data), o)
	err := Accept("nanaka", "fedi.moonchan.xyz", o)
	if err != nil {
		t.Error(err)
	}
}
