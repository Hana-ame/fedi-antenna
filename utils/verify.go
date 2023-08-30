package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-fed/httpsig"
	"github.com/iancoleman/orderedmap"
)

// httpsig
func RequestToAlgorithm(r *http.Request) httpsig.Algorithm {
	signature := r.Header.Get("Signature")
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

// keyId is something like "https://example.com/users/username#main-key", "" if failed
func RequestToPublicKeyId(r *http.Request) string {
	signature := r.Header.Get("Signature")
	for _, v := range strings.Split(signature, ",") {
		// log.Println(v)
		if strings.HasPrefix(v, "keyId") {
			// log.Println(v)
			keyId := strings.Split(v, "=")[1]
			keyId = keyId[1 : len(keyId)-1]
			return keyId
			// // 此处中断
			// publicKeyString := db.GetPublicKeyByKeyId(keyId)
			// //
			// publicKey, _ := mycrypto.ParsePublicKey(publicKeyString)
			// return publicKey
		}
	}
	return ""
}

func UserObjToPublicKeyPem(o *orderedmap.OrderedMap) string {
	if v, ok := o.Get("publicKey"); !ok {
		handleErr(errors.New("notfound: " + "publicKey"))
	} else {
		if pko, ok := v.(orderedmap.OrderedMap); !ok {
			handleErr(errors.New("typeerror: " + "publicKey"))
		} else {
			if pem, ok := pko.Get("publicKeyPem"); !ok {
				handleErr(errors.New("notfound: " + "publicKeyPem"))
			} else {
				return pem.(string)
			}
		}
	}
	return "NEVER"
}
