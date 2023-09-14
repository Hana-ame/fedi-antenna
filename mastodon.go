package main

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

func apiLookUp(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	c.JSON(400, gin.H{"error": "not supported"})
}

func apiAccount(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	c.JSON(400, gin.H{"error": "not supported"})
}

func apiAccountBlock(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	c.JSON(400, gin.H{"error": "not supported"})
}

func apiAccountUnBlock(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	c.JSON(400, gin.H{"error": "not supported"})
}

func apiAccountFollow(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	c.JSON(400, gin.H{"error": "not supported"})
}

func apiAccountUnFollow(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	c.JSON(400, gin.H{"error": "not supported"})
}

func apiAccountMute(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	c.JSON(400, gin.H{"error": "not supported"})
}

func apiAccountUnMute(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	c.JSON(400, gin.H{"error": "not supported"})
}
