package curl

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	tools "github.com/Hana-ame/fedi-antenna/Tools"
	mycurl "github.com/Hana-ame/fedi-antenna/Tools/my_curl"
	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
	"github.com/Hana-ame/fedi-antenna/httpsig"
)

// Decreped
func Get(id string) (*orderedmap.OrderedMap, error) {
	headers := mycurl.Headers{
		{"Accept", "application/ld+json"},
	}
	status, body, err := mycurl.Get("", headers, "", id)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("return http status %d", status)
	}
	o, err := tools.BytesToJson(body)
	return o, err
}

// privateKeyPem: -----BEGIN PUBLIC KEY-----
// MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsvkAzK1gZ4uz6m10I/N7
// XRNnE8p7/ix0gyXQGvvlQAK44TYgRaExi7aAJsNOzk71VkkrTxrUbRNBY1QmhVnP
// ytmgV3/1PZlbkYqzcRIl2KK0rPBfK72GMpynJp1v2BqlArVIVxIVq4SCjR7qsXVC
// 63C68K4Jt80IAerxHehZIhJQWN4pKAHkodZ8E0clPgdj9XJW5B6Qqd6zGMb9fai5
// LxBF4AIcDNg2oM+0Ckv35Fp8rEyJGW2ZEMsnIBrta2JmdRSieuQz6aICUhgkvddA
// mTtz4+wBFz/ARow/VmC6jO837XWzSzr3I2O+QErrTsbnpkGdU8qy5IgZ/V2PjRNG
// 1QIDAQAB
// -----END PUBLIC KEY-----
// publicKeyID: ttps://mstdn.jp/users/nanakananoka#main-key
func GetWithSignDefault(id string) (*orderedmap.OrderedMap, error) {
	pk, err := tools.ReadKeyFromFile(os.Getenv("USERNAME") + ".pem")
	if err != nil {
		return nil, err
	}
	pem, err := tools.MarshalPrivateKey(pk)
	if err != nil {
		return nil, err
	}
	keyID := "https://" + os.Getenv("HOST") + "/user/" + os.Getenv("USERNAME") + "#main-key"
	return GetWithSign(id, pem, keyID)
}

func GetWithSign(id string, privateKeyPem []byte, publicKeyID string) (*orderedmap.OrderedMap, error) {

	r, _ := http.NewRequest(http.MethodGet, id, nil)
	// 获取当前时间并格式化为 RFC1123 格式
	date := time.Now().UTC().Format(http.TimeFormat)
	r.Header.Set("Date", date)
	r.Header.Set("Accept", "application/ld+json")

	privateKey, _ := tools.ParsePrivateKey(privateKeyPem)

	digestAlgorithm := httpsig.DigestSha256
	// The "Date" and "Digest" headers must already be set on r, as well as r.URL.
	headersToSign := []string{httpsig.RequestTarget, "date"}
	signer, chosenAlgo, err := httpsig.NewSigner(nil, digestAlgorithm, headersToSign, httpsig.Signature, 3600)
	if err != nil {
		return nil, err
	}
	_ = chosenAlgo
	// To sign the digest, we need to give the signer a copy of the body...
	// ...but it is optional, no digest will be signed if given "nil"
	// body := ...
	// If r were a http.ResponseWriter, call SignResponse instead.

	err = signer.SignRequest(privateKey, publicKeyID, r, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return tools.ReaderToJson(resp.Body)
}

func PostWithSignDefault(id string, body []byte) (*orderedmap.OrderedMap, error) {
	pk, err := tools.ReadKeyFromFile(os.Getenv("USERNAME") + ".pem")
	if err != nil {
		return nil, err
	}
	pem, err := tools.MarshalPrivateKey(pk)
	if err != nil {
		return nil, err
	}
	keyID := "https://" + os.Getenv("HOST") + "/user/" + os.Getenv("USERNAME") + "#main-key"
	return PostWithSign(id, body, pem, keyID)
}

func PostWithSign(id string, body, privateKeyPem []byte, publicKeyID string) (*orderedmap.OrderedMap, error) {

	r, _ := http.NewRequest(http.MethodPost, id, bytes.NewReader(body))
	// 获取当前时间并格式化为 RFC1123 格式
	date := time.Now().UTC().Format(http.TimeFormat)
	r.Header.Set("Date", date)
	r.Header.Set("Accept", "application/ld+json")

	privateKey, _ := tools.ParsePrivateKey(privateKeyPem)

	digestAlgorithm := httpsig.DigestSha256
	// The "Date" and "Digest" headers must already be set on r, as well as r.URL.
	headersToSign := []string{httpsig.RequestTarget, "date", "digest"}
	signer, chosenAlgo, err := httpsig.NewSigner(nil, digestAlgorithm, headersToSign, httpsig.Signature, 3600)
	if err != nil {
		return nil, err
	}
	_ = chosenAlgo
	// To sign the digest, we need to give the signer a copy of the body...
	// ...but it is optional, no digest will be signed if given "nil"
	// body := ...
	// If r were a http.ResponseWriter, call SignResponse instead.

	err = signer.SignRequest(privateKey, publicKeyID, r, body)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return tools.ReaderToJson(resp.Body)
}
