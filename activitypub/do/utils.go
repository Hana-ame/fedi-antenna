package do

import (
	"bytes"
	"crypto"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Hana-ame/fedi-antenna/Tools/myfetch"
	"github.com/Hana-ame/fedi-antenna/activitypub/dao"
	"github.com/Hana-ame/fedi-antenna/activitypub/dao/model"
	"github.com/Hana-ame/fedi-antenna/utils"
)

func Fetch(method, url string, header map[string]string, body io.Reader) (*http.Response, error) {
	header = utils.MergeMaps(header, map[string]string{"Accept": "application/activity+json"})
	return myfetch.Fetch(method, url, header, body)
}

// owner is activitypubID with out '#main-key'
func FetchWithSign(
	owner string,
	method string, endpoint string, header map[string]string, body []byte,
) (
	*http.Response, error,
) {

	key := &model.PublicKey{
		Owner: owner,
	}
	err := dao.Read(key)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	pk, err := utils.ParsePrivateKey(key.PrivateKeyPem)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := fetchWithSign(
		pk, owner+"#main-key",
		method, endpoint, header, body,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, err
}

func fetchWithSign(
	privateKey crypto.PrivateKey, pubKeyID string,
	method, url string, header map[string]string, body []byte,
) (
	*http.Response, error,
) {
	r, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return nil, err
	}
	resp, err := myfetch.FetchWithRequest(r)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	return resp, nil
}
