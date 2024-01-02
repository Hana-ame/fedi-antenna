package actions

import (
	"bytes"
	"crypto"
	"io"
	"net/http"
	"time"

	"github.com/Hana-ame/fedi-antenna/Tools/myfetch"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

func Fetch(method, url string, header map[string]string, body io.Reader) (*http.Response, error) {
	header = utils.MergeMaps(header, map[string]string{"Accept": "application/activity+json"})
	return myfetch.Fetch(method, url, header, body)
}

// pubKeyOwner is ID with out '#main-key'
func FetchWithSign(
	pubKeyOwner string,
	method, url string, header map[string]string, body []byte,
) (
	*http.Response, error,
) {
	pk, err := dao.ReadPublicKeyByOwner(pubKeyOwner)
	if err != nil {
		return nil, err
	}

	return fetchWithSign(
		pk, pk.ID,
		method, url, header, body,
	)

}

func fetchWithSign(
	privateKey crypto.PrivateKey, pubKeyID string,
	method, url string, header map[string]string, body []byte,
) (
	*http.Response, error,
) {
	r, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	header = utils.MergeMaps(header, map[string]string{
		"host":         r.URL.Host,
		"date":         time.Now().UTC().Format(http.TimeFormat),
		"content-type": "application/activity+json",
	})
	for k, v := range header {
		r.Header.Set(k, v)
	}
	err = Sign(privateKey, pubKeyID, r, body)
	if err != nil {
		return nil, err
	}
	resp, err := myfetch.FetchWithRequest(r)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
