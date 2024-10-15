package main

import (
	"fmt"
	"net/http"
	"testing"

	tools "github.com/Hana-ame/fedi-antenna/Tools"
	mycurl "github.com/Hana-ame/fedi-antenna/Tools/my_curl"
	"github.com/Hana-ame/fedi-antenna/json-gold/ld"
)

// 现在Context在展开的时候挂了，理应有一个Context可以去overRide掉的。
// 现在开始trace

func TestCurl1(t *testing.T) {
	headers := mycurl.Headers{
		{"Accept", "application/json"},
	}
	code, body, err := mycurl.Curl(http.MethodGet, "", headers, "", "https://mstdn.jp/users/nanakananoka", nil)
	fmt.Println(err)
	// fmt.Println(string(body))
	fmt.Println(code)

	j, e := tools.BytesToJson(body)
	if e != nil {
		t.Error(e)
	}

	_ = j

	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	expanded, e := proc.Expand("https://mona.do/users/HaruUrara", options)
	if e != nil {
		t.Error(e)
	}

	ld.PrintDocument("JSON-LD expansion succeeded", expanded)

}

func TestXxx(t *testing.T) {
	var input any = map[string]interface{}{
		"@context":  "http://schema.org/",
		"@type":     "Person",
		"name":      "Jane Doe",
		"jobTitle":  "Professor",
		"telephone": "(425) 123-4567",
		"url":       "http://www.janedoe.com",
	}
	iri, isString := input.(string)
	fmt.Println(iri)
	fmt.Println(isString)
}
