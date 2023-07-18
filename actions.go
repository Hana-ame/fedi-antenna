package main

import (
	"bytes"
	"crypto"
	"log"
	"net/http"
	"time"

	"github.com/Hana-ame/fedi-antenna/testpkg"
	"github.com/go-fed/httpsig"
)

func SomeFun() {
	log.Println(testpkg.Test)
}

func NewSingedRequest(
	privateKey crypto.PrivateKey, pubKeyId string,
	method, url string, body []byte,
) (r *http.Request, err error) {
	r, err = http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return
	}
	r.Header.Set("host", r.URL.Host)
	r.Header.Set("date", time.Now().UTC().Format(http.TimeFormat))
	r.Header.Set("content-type", "application/activity+json")

	err = sign(privateKey, pubKeyId, r, body)
	if err != nil {
		return
	}

	return
}

func sign(privateKey crypto.PrivateKey, pubKeyId string, r *http.Request, body []byte) error {
	prefs := []httpsig.Algorithm{httpsig.RSA_SHA256}
	digestAlgorithm := httpsig.DigestSha256
	// The "Date" and "Digest" headers must already be set on r, as well as r.URL.
	headersToSign := []string{httpsig.RequestTarget, "host", "date", "digest", "content-type"}
	signer, chosenAlgo, err := httpsig.NewSigner(prefs, digestAlgorithm, headersToSign, httpsig.Signature, 1<<16)
	log.Println(chosenAlgo)
	if err != nil {
		return err
	}
	// To sign the digest, we need to give the signer a copy of the body...
	// ...but it is optional, no digest will be signed if given "nil"
	// body := []byte{}
	log.Println(string(body))

	// If r were a http.ResponseWriter, call SignResponse instead.
	return signer.SignRequest(privateKey, pubKeyId, r, body)
}
