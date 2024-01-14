package controller

import (
	"github.com/gin-gonic/gin"
)

func RegisterPath(r *gin.Engine) {
	// s2s send
	r.GET("/users/:name", Users)
	// r.GET("/users/:name/outbox", UsersOutbox)
	// r.GET("/users/:name/followers", UsersFollowers)
	// r.GET("/users/:name/following", UsersFollowing)
	// r.GET("/users/:name/collections/featured", UsersCollectionsFeatured)
	// r.GET("/users/:name/collections/tags", UsersCollectionsTags)
	// // https://mstdn.jp/users/meromero/statuses/110734394957749061
	// r.GET("/users/:name/statuses/:id", UsersStatuses)
	// // s2s recv
	r.POST("/users/:name/inbox", UsersInbox) // inbox, only on POST?
	r.POST("/inbox", Inbox)                  // inbox
}
