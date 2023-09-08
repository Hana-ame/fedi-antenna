package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/db"
	"github.com/Hana-ame/fedi-antenna/utils"
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/orderedmap"
)

var pureW = regexp.MustCompile(`^\w+$`)

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

// todo
// notwork
func apiFollow(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	c.JSON(400, gin.H{"error": "not supported"})
}

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
	// 判断是否存在 not used
	// if !core.IsUserExist(username, host) {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	// 	return
	// }

	// 形成返回数据
	webfingerObj, err := core.WebfingerObj(username, host)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, webfingerObj)
}

// local api

// users
// POST, GET, PUT
// TODO: move into core

// create new user, post an object that refers a user
func apiPostUser(c *gin.Context) {
	var user db.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check
	if !pureW.MatchString(user.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user name illegel"})
		return
	}

	// create user
	if err := core.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// err := os.WriteFile(utils.ParseNameAndHostToAcctStr(user.Username, user.Host)+".pem", pkPem, 0644)
	// if err != nil {
	// 	fmt.Println("Error writing file:", err)
	// 	return
	// }

	c.JSON(http.StatusOK, user)
}

// return a user, get an object thar refers a user
func apiGetUser(c *gin.Context) {
	name := c.Param("name")
	host := c.Request.Host

	user, err := db.ReadUserByKey(name, host)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// delete a user
func apiDeleteUser(c *gin.Context) {
	name := c.Param("name")
	host := c.Request.Host

	user := &db.User{Name: name, Host: host}
	err := core.DeleteUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// renew a user, put an object that refers a user
func apiPutUser(c *gin.Context) {
	// name := c.Param("name")
	// host := c.Request.Host
	// TODO: verify

	var user db.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Query a remote user

func apiQueryRemoteUser(c *gin.Context) {
	acct := c.Query("acct")

	user, err := core.GetRemoteUser(acct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, user.O) // return text/plain; charset=utf-8

}

// get something like {"action":"follow", actor:"me@local.host", object:"you@yours.domain"}
func apiAction(c *gin.Context) {
	var o core.ApiAction
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO
	// core.HandleAction(o)

}
