package curl

import (
	"fmt"
	"testing"

	tools "github.com/Hana-ame/fedi-antenna/Tools"
	"github.com/Hana-ame/fedi-antenna/Tools/debug"
)

func TestActivityPub(t *testing.T) {
	o, e := Get("https://mstdn.jp/users/nanakananoka")
	fmt.Println(e)
	fmt.Println(o)
	debug.OrderedMap(o)
}

func TestFollowers(t *testing.T) {
	o, e := Get("https://mstdn.jp/users/nanakananoka/followers")
	fmt.Println(e)
	fmt.Println(o)
	debug.OrderedMap(o)
}

func TestGetWithSign(t *testing.T) {
	pk, _ := tools.GeneratePrivateKey()
	pem, _ := tools.MarshalPrivateKey(pk)
	o, e := GetWithSign("https://getip.moonchan.xyz/echo.json", pem, "pubkeyID")
	fmt.Println(e)
	fmt.Println(o)
	debug.OrderedMap(o)
}

// not tested
func TestPostWithSign(t *testing.T) {
	pk, _ := tools.GeneratePrivateKey()
	pem, _ := tools.MarshalPrivateKey(pk)
	o, e := PostWithSign("https://getip.moonchan.xyz/echo", []byte("123123123"), pem, "pubkeyID")
	fmt.Println(e)
	fmt.Println(o)
	debug.OrderedMap(o)
}
