package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"
)

var jar, _ = cookiejar.New(nil)

// var client = &http.Client{
// 	CheckRedirect: func(req *http.Request, via []*http.Request) error {
// 		return http.ErrUseLastResponse
// 	},
// 	Jar: jar,
// }

var client = func() *http.Client {
	proxyURL, _ := url.Parse("http://localhost:10809")
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	return &http.Client{
		Transport: tr,
		Jar:       jar,
	}
}()

func printResp(resp *http.Response) {
	respText, _ := io.ReadAll(resp.Body)
	log.Println(string(respText))
}

//	func TestPkg(t *testing.T) {
//		log.Println(testpkg.test)
//	}
func TestMstdnJP(t *testing.T) {
	o := genFollowObj("https://fedi.moonchan.xyz/users/nanaka", "https://mstdn.jp/users/meromero")
	j, _ := json.Marshal(o)
	pk, _ := tools.ReadKeyFromFile("nanaka.pem")
	// r, _ := NewSingedRequest(pk, "https://fedi.moonchan.xyz/users/nanaka#main-key", "POST", "https://moonchan.xyz/api-pack/echo", j)
	r, _ := NewSingedRequest(pk, "https://fedi.moonchan.xyz/users/nanaka#main-key", "POST", "https://mstdn.jp/users/meromero/inbox", j)
	log.Println(r.Header)
	resp, err := client.Do(r)
	log.Println(err)
	printResp(resp)
}

// passed?
func TestReq(t *testing.T) {
	pk, err := tools.ReadKeyFromFile("nanaka.pem")
	log.Println(err)
	r, err := NewSingedRequest(pk, "id", "POST", "https://moonchan.xyz/api-pack/echo", []byte("something else"))
	log.Println(err)
	resp, err := client.Do(r)
	log.Println(err)
	printResp(resp)
}
