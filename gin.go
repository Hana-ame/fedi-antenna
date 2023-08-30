package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/utils"
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/orderedmap"
)

func users(c *gin.Context) {
	name := c.Param("name")
	if name != "nanaka" {
		c.JSON(http.StatusNotFound, gin.H{"error": "only support nanaka@fedi.moonchan.xyz"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"1": "2"})
}

// not work
func usersOutbox(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	verify(c.Request)

	c.JSON(400, gin.H{"error": "not supported"})

}

// not work
func usersFollowers(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	verify(c.Request)
	c.JSON(400, gin.H{"error": "not supported"})

}

// not work
func usersFollowing(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	verify(c.Request)
	c.JSON(400, gin.H{"error": "not supported"})
}

// not work
func usersCollectionsFeatured(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	verify(c.Request)
	c.JSON(400, gin.H{"error": "not supported"})
}

// not work
func usersCollectionsTags(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	verify(c.Request)
	c.JSON(400, gin.H{"error": "not supported"})
}

// notwork
func usersStatuses(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	c.JSON(400, gin.H{"error": "not supported"})
}

// notwork
func apiFollow(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	c.JSON(400, gin.H{"error": "not supported"})
}

// inbox

func usersInbox(c *gin.Context) {
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

	c.JSON(400, gin.H{"error": "not supported"})

}

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

// webfinger 的 server
func webfinger(c *gin.Context) {
	// 被访问
	// 取得query
	resource := c.Query("resource")
	if !strings.HasPrefix(resource, "acct:") {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "format error"})
		return
	}
	// 取出成为 uesrname 和 host
	acct := strings.TrimPrefix(resource, "acct:")
	username, host := utils.ParseAcctStrToUserAndHost(acct)
	// 判断是否存在
	if !core.IsUserExist(username, host) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	// 形成返回数据
	webfingerObj, err := core.CreateWebfingerObj(username, host)
	if err != nil { // 这里会跑到么。
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, webfingerObj)
}
