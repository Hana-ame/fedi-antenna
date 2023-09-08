package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/httpsig"
	"github.com/Hana-ame/fedi-antenna/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// no use
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// webfinger
	r.GET(".well-known/webfinger", webfinger)
	// s2s send
	r.GET("/users/:name", users)
	r.GET("/users/:name/outbox", usersOutbox)
	r.GET("/users/:name/followers", usersFollowers)
	r.GET("/users/:name/following", usersFollowing)
	r.GET("/users/:name/collections/featured", usersCollectionsFeatured)
	r.GET("/users/:name/collections/tags", usersCollectionsTags)
	// https://mstdn.jp/users/meromero/statuses/110734394957749061
	r.GET("/users/:name/statuses/:id", usersStatuses)
	// s2s recv
	r.POST("/users/:name/inbox", usersInbox) // inbox, only on POST?
	r.POST("/inbox", inbox)                  // inbox
	// users
	r.POST("/api/v1/users", apiPostUser)
	r.GET("/api/v1/users/:name", apiGetUser)
	r.PUT("/api/v1/users/:name", apiPutUser)
	r.DELETE("/api/v1/users/:name", apiDeleteUser)

	r.GET("/api/v1/remoteusers", apiQueryRemoteUser)

	r.POST("/api/v1/action", apiAction)

	r.Run("0.0.0.0:3000") // listen and serve on 0.0.0.0:3000
}

// utils
// 这个需要之后再看一下修改一下
func verify(r *http.Request) (err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("%s", e)
		}
	}()

	algo := utils.RequestToAlgorithm(r)
	pubKeyId := utils.RequestToPublicKeyId(r)
	userObj, err := utils.FetchObj("GET", pubKeyId, nil) // this.
	if err != nil {
		panic(err)
	}
	pubKeyPem := utils.UserObjToPublicKeyPem(userObj)
	pubKey, err := utils.ParsePublicKey(pubKeyPem)
	if err != nil {
		panic(err)
	}
	err = httpsig.Verify(r, algo, pubKey)
	log.Println(err)
	return err
}
