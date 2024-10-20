package main

import (
	"crypto"
	"crypto/rsa"
	"fmt"
	"net/http"
	"strings"

	"github.com/Hana-ame/fedi-antenna/Tools/debug"
	"github.com/Hana-ame/fedi-antenna/httpsig"
	"github.com/Hana-ame/fedi-antenna/scheduler"
)

// usage(gin):
// verify(c.Request)
func verify(r *http.Request) error {
	verifier, err := httpsig.NewVerifier(r)
	if err != nil {
		return err
	}
	// pubKeyId := verifier.KeyId()
	var algo httpsig.Algorithm = parseAlgorithm(r.Header.Get("Signature"))
	var pubKey crypto.PublicKey = parsePublicKey(r.Header.Get("Signature"))
	// The verifier will verify the Digest in addition to the HTTP signature
	// log.Println(pubKey, algo)

	return verifier.Verify(pubKey, algo)
}

func parseAlgorithm(signature string) httpsig.Algorithm {
	for _, v := range strings.Split(signature, ",") {
		// log.Println(v)
		if strings.HasPrefix(v, "algorithm") {
			// log.Println(v)
			algorithm := strings.Split(v, "=")[1]
			algorithm = algorithm[1 : len(algorithm)-1]
			return httpsig.Algorithm(algorithm)
		}
	}
	return ""
}

func parsePublicKey(signature string) *rsa.PublicKey {
	defer func() {
		if e := recover(); e != nil {
			debug.E("parsePublicKey", fmt.Errorf("%v", e))
			return
		}
	}()
	for _, v := range strings.Split(signature, ",") {
		// log.Println(v)
		if strings.HasPrefix(v, "keyId") {
			// log.Println(v)
			keyID := strings.Split(v, "=")[1]
			keyID = keyID[1 : len(keyID)-1]
			ID := strings.Split(keyID, "#")[0]
			// 此处中断
			// publicKeyString := fetchPublicKeyByKeyId(keyID)
			//
			apUser, _ := scheduler.ByApID(ID)
			apPbulicKey, ok := apUser.Get("PublicKey")
			if !ok {
				return nil
			}
			publicKey := apPbulicKey.Get("TODO")
			return publicKey
		}
	}
	return nil
}

func fetchPublicKeyByKeyId(keyId string) string {
	return "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtYjfX1ONI/uHwqMMbfYt\n568/VjE/C0I/7V1SfRoX1cUWc0H7XE2hJz1DoDoLsKC4pMIjeWLlr49L2liiJYzF\nSCIJBVzAWsl7aiqIfKPKH9wS07lrCp7iz2B0nfvG1EoskncHe1H2LqWpOEUZirYo\nyFFJKM+xYlxqzLuwcw3NHq2Mf149iCeOAB4ZKZVBq8R8arb6MtFepifnWFz4Hf4I\nep8OujJ2SYRlv9+li//HqR/PwbYmkfvTi+64F+/JiTmHXzGDBzGdLbUhU+ZuZN9b\n6aZdruRsIVeeRvqGa2XW6oDcVadNXK6ovVY29TjimaCYgZfdTGpfDiBWAQQITjfs\nVQIDAQAB\n-----END PUBLIC KEY-----\n"
}
