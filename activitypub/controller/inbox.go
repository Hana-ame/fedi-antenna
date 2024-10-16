package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
	"github.com/Hana-ame/fedi-antenna/activitypub/handler"
	"github.com/gin-gonic/gin"
)

func UsersInbox(c *gin.Context) {
	name := c.Param("name")
	host := c.GetHeader("Host")

	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	o := orderedmap.New()
	err = json.Unmarshal(b, &o)
	if err != nil {
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

	err = verify(c, b)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	err = handler.Inbox(o, name, host, err)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.String(http.StatusOK, "")
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

	err = verify(c, b)
	if err != nil {
		log.Println(err)
		// c.JSON(http.StatusUnauthorized, err.Error())
		// return
	}
	if err := handler.Inbox(o, name, host, err); err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.String(http.StatusOK, "")
}
