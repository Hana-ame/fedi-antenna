package main

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/iancoleman/orderedmap"
)

var host = "fedi.moonchan.xyz"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET(WEBFINGER_PATH, webfinger)
	r.GET("/users/:name", users)
	r.POST("/users/:name/inbox", usersInbox)
	r.POST("/inbox", inbox)
	r.Run("0.0.0.0:3000") // listen and serve on 0.0.0.0:3000
}
func usersInbox(c *gin.Context) {
	o := orderedmap.New()
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	log.Println(verify(c.Request))

	b, _ = json.Marshal(o)
	log.Println(string(b))
	// Do something with the JSON data
	// For example, print the received data
	// c.JSON(200, o)
	c.JSON(400, gin.H{"error": "not supported"})

}

func inbox(c *gin.Context) {
	o := orderedmap.New()
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	b, _ = json.Marshal(o)
	log.Println(string(b))
	// Do something with the JSON data
	// For example, print the received data
	// c.JSON(200, o)
	c.JSON(400, gin.H{"error": "not supported"})

}
