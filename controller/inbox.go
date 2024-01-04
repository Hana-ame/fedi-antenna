package controller

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/activitypub/handler"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/orderedmap"
	"github.com/gin-gonic/gin"
)

func UsersInbox(c *gin.Context) {
	name := c.Param("name")

	o := orderedmap.New()
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	b, _ = json.Marshal(o)
	log.Println(string(b))

	// verify
	if err := Verify(c.Request); err != nil {
		log.Println(err)
		return
	}
	_, digest := utils.ParseDigest(c.GetHeader("Digest"))
	sha256 := sha256.Sum256([]byte(b))
	encoded := base64.StdEncoding.EncodeToString([]byte(sha256[:]))
	if digest != encoded {
		log.Println("digest != encoded")
		return
	}
	// end of verify

	handler.Inbox(b, name)

	c.JSON(http.StatusAccepted, gin.H{"error": "StatusNotImplemented"})

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

	if err := Verify(c.Request); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusNotImplemented, gin.H{"error": "not supported"})

}
