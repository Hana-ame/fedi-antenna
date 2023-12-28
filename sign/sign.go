package sign

import (
	"crypto"
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/go-fed/httpsig"
)

func Verify(r *http.Request) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = (fmt.Errorf("%s", e))
		}
	}()
	verifier, err := httpsig.NewVerifier(r)
	if err != nil {
		return err
	}
	var algo httpsig.Algorithm = parseAlgorithm(r.Header.Get("Signature"))
	var pubKey crypto.PublicKey = parsePublicKey(r.Header.Get("Signature"))

	return verifier.Verify(pubKey, algo)
}

func parseAlgorithm(signature string) httpsig.Algorithm {
	for _, v := range strings.Split(signature, ",") {
		if strings.HasPrefix(v, "algorithm") {
			algorithm := strings.TrimPrefix(v, "algorithm=")
			algorithm = strings.TrimPrefix(algorithm, "\"")
			algorithm = strings.TrimSuffix(algorithm, "\"")
			return httpsig.Algorithm(algorithm)
		}
	}
	return ""
}

func parsePublicKey(signature string) *rsa.PublicKey {
	for _, v := range strings.Split(signature, ",") {
		if strings.HasPrefix(v, "keyId") {
			keyId := strings.TrimPrefix(v, "keyId=")
			keyId = strings.TrimPrefix(keyId, "\"")
			keyId = strings.TrimSuffix(keyId, "\"")
			keyId = strings.TrimSuffix(keyId, "#main-key")

			pk, err := core.ReadPublicKeyByOwner(keyId)
			if err != nil {
				log.Println(err)
				return nil
			}

			publicKey, _ := utils.ParsePublicKey(pk.PublicKeyPem)
			return publicKey
		}
	}
	return nil
}
