package controller

import (
	"github.com/gin-gonic/gin"
)

func RegisterPath(r *gin.Engine) {
	r.GET("/users/:name", Users)
	// r.GET("/users/:name/outbox", UsersOutbox)
	// r.GET("/users/:name/followers", UsersFollowers)
	// r.GET("/users/:name/following", UsersFollowing)
	// r.GET("/users/:name/collections/featured", UsersCollectionsFeatured)
	// r.GET("/users/:name/collections/tags", UsersCollectionsTags)
	r.GET("/users/:name/statuses/:id", UsersStatuses)

	r.POST("/users/:name/inbox", UsersInbox) // inbox, only on POST?
	r.POST("/inbox", Inbox)                  // inbox
}
