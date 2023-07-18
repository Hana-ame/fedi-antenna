package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iancoleman/orderedmap"
)

const WEBFINGER_PATH = ".well-known/webfinger"

// .well-known/webfinger?resource=acct:nanaka@fedi.moonchan.xyz
func webfinger(c *gin.Context) {
	resource := c.Query("resource")
	if resource != "acct:nanaka@fedi.moonchan.xyz" {
		c.JSON(http.StatusNotFound, gin.H{"error": "only support nanaka@fedi.moonchan.xyz"})
		return
	}
	o := newWebfingerSubjectObj(resource)
	c.JSON(http.StatusOK, o)
}

func newWebfingerSubjectObj(subject string) *orderedmap.OrderedMap {
	log.Println(subject)
	name := strings.Split(subject, ":")[1]
	name = strings.Split(name, "@")[0]
	o := orderedmap.New()
	o.Set("subject", subject)
	o.Set("aliases", []string{"https://" + host + "/@" + name, "https://" + host + "/users/" + name})
	o.Set("links", []*orderedmap.OrderedMap{
		newLinksObj("http://webfinger.net/rel/profile-page", "text/html", "https://"+host+"/@"+name, ""),
		newLinksObj("self", "application/activity+json", "https://"+host+"/users/"+name, ""),
		newLinksObj("http://ostatus.org/schema/1.0/subscribe", "", "", "https://"+host+"/authorize_interaction?uri={uri}"),
	})

	return o
}

func newLinksObj(relStr, typeStr, hrefStr, templateStr string) *orderedmap.OrderedMap {
	o := orderedmap.New()
	o.Set("rel", relStr)
	if typeStr != "" {
		o.Set("type", typeStr)
	}
	if hrefStr != "" {
		o.Set("href", hrefStr)
	}
	if templateStr != "" {
		o.Set("template", templateStr)
	}
	return o
}
