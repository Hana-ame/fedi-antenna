package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/orderedmap"
)

// inbox

// todo
func usersInbox(c *gin.Context) {
	o := orderedmap.New()
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	verify := verify(c.Request)
	// var verify error = nil

	core.Inbox(c.Request.Header, o, verify)

	// bh, _ := json.Marshal(c.Request.Header)
	// log.Println(string(bh))

	bb, _ := json.Marshal(o)
	log.Println(string(bb))

	c.JSON(400, gin.H{"error": "not supported"})

}

// todo
func inbox(c *gin.Context) {
	o := orderedmap.New()
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// verify := verify(c.Request)
	var verify error = nil

	core.Inbox(c.Request.Header, o, verify)

	// bh, _ := json.Marshal(c.Request.Header)
	// log.Println(string(bh))

	bb, _ := json.Marshal(o)
	log.Println(string(bb))

	// Do something with the JSON data
	// For example, print the received data
	// c.JSON(200, o)
	c.JSON(400, gin.H{"error": "not supported"})

}

// only test
// /users/:name
func users(c *gin.Context) {
	name := c.Param("name")
	host := c.Request.Host
	o, err := core.UserObj(name, host)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, o)
}

// todo
// not work
func usersOutbox(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	verify(c.Request)

	c.JSON(400, gin.H{"error": "not supported"})

}

// todo
// not work
func usersFollowers(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	verify(c.Request)
	c.JSON(400, gin.H{"error": "not supported"})

}

// todo
// not work
func usersFollowing(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	verify(c.Request)
	c.JSON(400, gin.H{"error": "not supported"})
}

// todo
// not work
func usersCollectionsFeatured(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	verify(c.Request)
	c.JSON(400, gin.H{"error": "not supported"})
}

// todo
// not work
func usersCollectionsTags(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	verify(c.Request)
	c.JSON(400, gin.H{"error": "not supported"})
}

// todo
// notwork
func usersStatuses(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	c.JSON(400, gin.H{"error": "not supported"})
}
