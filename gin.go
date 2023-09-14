package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/db"
	"github.com/gin-gonic/gin"
)

var pureW = regexp.MustCompile(`^\w+$`)

// todo
// notwork
func apiFollow(c *gin.Context) {
	b, _ := json.Marshal(c.Request.Header)
	log.Println(string(b))

	c.JSON(400, gin.H{"error": "not supported"})
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
