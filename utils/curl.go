package utils

import (
	"bytes"
	"crypto"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/Hana-ame/fedi-antenna/httpsig"
	"github.com/iancoleman/orderedmap"
)

var jar, _ = cookiejar.New(nil)
var proxyURL, _ = url.Parse("http://localhost:10809")

// use this in formal version
// var client = &http.Client{
// // add this to disable redirect, do not use it
// 	CheckRedirect: func(req *http.Request, via []*http.Request) error {
// 		return http.ErrUseLastResponse
// 	},
// 	Jar: jar,
// }

var client = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	},
	Jar: jar,
}

func Do(req *http.Request) (*http.Response, error) {
	return client.Do(req)
}

func NewSingedRequest(
	privateKey crypto.PrivateKey, pubKeyId string,
	method, url string, body []byte,
) (r *http.Request, err error) {
	r, err = http.NewRequest(method, url, bytes.NewReader(body))
	handleErr(err)

	r.Header.Set("host", r.URL.Host)
	r.Header.Set("date", time.Now().UTC().Format(http.TimeFormat))
	r.Header.Set("content-type", "application/activity+json")

	err = httpsig.Sign(privateKey, pubKeyId, r, body)
	handleErr(err)

	return
}

func FetchObj(method, url string, body io.Reader, headers ...map[string][]string) (o *orderedmap.OrderedMap, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%s", e)
		}
	}()

	r, err := http.NewRequest(method, url, body)
	handleErr(err)

	r.Header.Set("Accept", "application/ld+json")
	r.Header.Set("Accept", "application/activity+json")

	for _, h := range headers {
		for k, va := range h {
			for _, v := range va {
				r.Header.Set(k, v)
			}
		}
	}

	resp, err := Do(r)
	handleErr(err)

	// need review
	respData, err := io.ReadAll(resp.Body)
	handleErr(err)

	o = orderedmap.New()
	err = json.Unmarshal(respData, &o)
	handleErr(err)

	return
}
