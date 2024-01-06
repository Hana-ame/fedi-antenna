package controller

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/activitypub/handler"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/orderedmap"
	"github.com/gin-gonic/gin"
)

func UsersInbox(c *gin.Context) {
	Inbox(c)
	// name := c.Param("name")
	// host := c.GetHeader("Host")

	// o := orderedmap.New()
	// if err := c.ShouldBindJSON(&o); err != nil {
	// 	log.Println(err)
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }

	// b, _ := json.Marshal(c.Request.Header)
	// log.Println(string(b))

	// b, _ = json.Marshal(o)
	// log.Println(string(b))

	// err := verify(c, b)
	// if err != nil {
	// 	log.Println(err)
	// 	c.JSON(http.StatusUnauthorized, err.Error())
	// 	return
	// }

	// handler.Inbox(o, name, host, err)

	// c.JSON(http.StatusAccepted, gin.H{"error": "StatusNotImplemented"})
}

func Inbox(c *gin.Context) {
	name := c.Param("name")
	host := c.GetHeader("Host")

	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	o := orderedmap.New()
	if err := json.Unmarshal(b, &o); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// for debug
	log.Println(string(b))
	h, _ := json.Marshal(c.Request.Header)
	log.Println(string(h))
	j, _ := json.Marshal(o)
	log.Println(string(j))

	if err := verify(c, b); err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	if err := handler.Inbox(o, name, host, err); err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

func verify(c *gin.Context, body []byte) error {
	// verify
	if err := Verify(c.Request); err != nil {
		log.Println(err)
		return err
	}
	_, digest := utils.ParseDigest(c.GetHeader("Digest"))
	sha256 := sha256.Sum256([]byte(body))
	encoded := base64.StdEncoding.EncodeToString([]byte(sha256[:]))
	if digest != encoded {
		log.Printf("digest != encoded\n")
		return fmt.Errorf("digest != encoded")
	}
	return nil
}
