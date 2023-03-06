package webfinger

import (
	"fmt"
)

const WebFingerPath = "/.well-known/webfinger"

var domain = "v.meromeromeiro.top"

// var checkExistCallback func(username string) bool = func(username string) bool { return true }

func SetDomain(_domain string) {
	domain = _domain
}

func CheckDomain(_domain string) bool {
	return domain == _domain
}

type Link struct {
	HRef       string             `json:"href,omitempty"`
	Type       string             `json:"type,omitempty"`
	Rel        string             `json:"rel,omitempty"`
	Template   string             `json:"template,omitempty"`
	Properties map[string]*string `json:"properties,omitempty"`
	Titles     map[string]string  `json:"titles,omitempty"`
}

type Resource struct {
	Subject    string            `json:"subject,omitempty"`
	Aliases    []string          `json:"aliases,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
	Links      []Link            `json:"links"`
}

func GetResource(username string, domain string) *Resource {
	// return mockGet(username)

	// TODO: sould connect to database to find Aliases
	self := &Link{
		Rel:  "self",
		Type: "application/activity+json",
		HRef: fmt.Sprintf("https://%s/users/%s", domain, username),
	}
	// porfile
	profilePage := &Link{
		Rel:  "http://webfinger.net/rel/profile-page",
		Type: "text/html",
		HRef: fmt.Sprintf("https://%s/@%s", domain, username),
	}
	// what's this used for?
	subscribe := &Link{
		Rel:      "http://ostatus.org/schema/1.0/subscribe",
		Template: fmt.Sprintf("https://%s/authorize-follow?acct={uri}", domain),
	}

	r := &Resource{
		Subject: fmt.Sprintf("acct:%s@%s", username, domain),
		Aliases: nil, // TODO
		Links: []Link{
			*self,
			*profilePage,
			*subscribe,
		},
	}

	// log.Println(r)

	return r
	// return mockGet(username)
}
