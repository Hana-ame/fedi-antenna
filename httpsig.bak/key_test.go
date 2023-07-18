package httpsig

import (
	"fmt"
	"testing"
)

func getChecker(t *testing.T) func(err error) {
	return func(err error) {
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestXX(t *testing.T) {
	key := GenerateKey()
	a := MarshalPrivateKey(key)
	fmt.Println(string(a))
}

func TestXXX(t *testing.T) {
	check := getChecker(t)

	k, err := ReadKeyFromFile("privateKey.pem")
	check(err)

	s := MarshalPublicKey(&k.PublicKey)

	fmt.Println(string(s))

}
