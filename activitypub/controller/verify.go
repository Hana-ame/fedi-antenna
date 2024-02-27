package controller

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/httpsig"
	"github.com/gin-gonic/gin"
)

func verify(c *gin.Context, body []byte) error {
	// verify
	if err := Verify(c.Request); err != nil {
		log.Println(err)
		return err
	}
	_, digest := parseDigest(c.GetHeader("Digest"))
	sha256 := sha256.Sum256([]byte(body))
	encoded := base64.StdEncoding.EncodeToString([]byte(sha256[:]))
	if digest != encoded {
		log.Printf("digest != encoded\n")
		return fmt.Errorf("digest != encoded")
	}
	return nil
}

// signature
// "SHA-256=8RIlimPwETDMkWQMI59d0gm9dqhzKGtX0CsEcahxxOE=" => "SHA-256", "8RIlimPwETDMkWQMI59d0gm9dqhzKGtX0CsEcahxxOE="
func parseDigest(d string) (algorithm, digest string) {
	arr := strings.SplitN(d, "=", 2)
	if len(arr) != 2 {
		return
	}
	return arr[0], arr[1]
}

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
