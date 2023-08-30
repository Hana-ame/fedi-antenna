package httpsig

import (
	"crypto"
	"crypto/rsa"
	"net/http"

	"github.com/go-fed/httpsig"
)

// usage(in gin):
// verify(c.Request)
func Verify(r *http.Request, algo httpsig.Algorithm, pubKey *rsa.PublicKey) error {
	verifier, err := httpsig.NewVerifier(r)
	if err != nil {
		return err
	}
	// pubKeyId := verifier.KeyId()
	// var algo httpsig.Algorithm = parseAlgorithm(r.Header.Get("Signature"))
	// var pubKey crypto.PublicKey = parsePublicKey(r.Header.Get("Signature"))
	// The verifier will verify the Digest in addition to the HTTP signature
	// log.Println(pubKey, algo)

	return verifier.Verify(pubKey, algo)
}

// usage:
// fill the inputs
func Sign(privateKey crypto.PrivateKey, pubKeyId string, r *http.Request, body []byte) error {
	prefs := []httpsig.Algorithm{httpsig.RSA_SHA256}
	digestAlgorithm := httpsig.DigestSha256
	// The "Date" and "Digest" headers must already be set on r, as well as r.URL.
	headersToSign := []string{httpsig.RequestTarget, "host", "date", "digest", "content-type"}
	signer, chosenAlgo, err := httpsig.NewSigner(prefs, digestAlgorithm, headersToSign, httpsig.Signature, 1<<16)
	// log.Println(chosenAlgo)
	_ = chosenAlgo
	if err != nil {
		return err
	}
	// To sign the digest, we need to give the signer a copy of the body...
	// ...but it is optional, no digest will be signed if given "nil"
	// body := []byte{}
	// log.Println(string(body))

	// If r were a http.ResponseWriter, call SignResponse instead.
	return signer.SignRequest(privateKey, pubKeyId, r, body)
}
