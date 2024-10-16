package controller

import (
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Hana-ame/fedi-antenna/activitypub/dao"
	"github.com/Hana-ame/fedi-antenna/activitypub/dao/model"
	"github.com/Hana-ame/fedi-antenna/activitypub/do"
	"github.com/Hana-ame/fedi-antenna/utils"
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
	algo := parseAlgorithm(r.Header.Get("Signature"))
	pubKey := parsePublicKey(r.Header.Get("Signature"))

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
	log.Println("algorithm not found")
	return ""
}

func parsePublicKey(signature string) *rsa.PublicKey {
	for _, v := range strings.Split(signature, ",") {
		if strings.HasPrefix(v, "keyId") {
			keyId := strings.TrimPrefix(v, "keyId=")
			keyId = strings.TrimPrefix(keyId, "\"")
			keyId = strings.TrimSuffix(keyId, "\"")
			owner := strings.TrimSuffix(keyId, "#main-key")

			pk := &model.PublicKey{Owner: owner}
			err := dao.Read(pk)
			if err != nil {
				// TBD: not checked
				_, err = do.User(owner)
				if err != nil {
					log.Println(err)
					return nil
				}
				err = dao.Read(pk)
				if err != nil {
					log.Println(err)
					return nil
				}
			}

			publicKey, err := utils.ParsePublicKey(pk.PublicKeyPem)
			if err != nil {
				log.Println(err)
				return nil
			}
			return publicKey
		}
	}
	log.Println("keyId not found")
	return nil
}
