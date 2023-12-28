package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/core/sign"
	"github.com/Hana-ame/orderedmap"
	"github.com/gin-gonic/gin"
)

func UsersInbox(c *gin.Context) {
	o := orderedmap.New()
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	b, _ = json.Marshal(o)
	log.Println(string(b))

	if err := sign.Verify(c.Request); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusNotImplemented, gin.H{"error": "StatusNotImplemented"})

}

func Inbox(c *gin.Context) {
	o := orderedmap.New()
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	b, _ = json.Marshal(o)
	log.Println(string(b))

	if err := sign.Verify(c.Request); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusNotImplemented, gin.H{"error": "not supported"})

}
