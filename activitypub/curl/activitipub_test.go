package curl

import (
	"fmt"
	"os"
	"testing"

	tools "github.com/Hana-ame/fedi-antenna/Tools"
	"github.com/Hana-ame/fedi-antenna/Tools/debug"
)

// func TestMain(m *testing.M) {

// }

const PEM_FILE = "pk.pem"
const pem = `-----BEGIN PRIVATE KEY-----
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
7ZMDJORAPKNrD/OnHs8RjMjzg13N
-----END PRIVATE KEY-----
`

func TestPem(t *testing.T) {

	pk, _ := tools.GeneratePrivateKey()
	pem, _ := tools.MarshalPrivateKey(pk)

	// 使用 ioutil.WriteFile 写入文件
	err := os.WriteFile(PEM_FILE, pem, 0644)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

}

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
	pem, _ := os.ReadFile(PEM_FILE)
	o, e := GetWithSign("https://getip.moonchan.xyz/echo.json", []byte(pem), "pubkeyID")
	fmt.Println(e)
	fmt.Println(o)
	debug.OrderedMap(o)
}

// not tested
func TestPostWithSign(t *testing.T) {
	// pem, _ := os.ReadFile(PEM_FILE)
	o, e := PostWithSign("https://getip.moonchan.xyz/echo", []byte("123123123"), []byte(pem), "pubkeyID")
	fmt.Println(e)
	fmt.Println(o)
	debug.OrderedMap(o)
}
