package activitypub

import (
	"github.com/Hana-ame/fedi-antenna/activitypub/controller"
	"github.com/gin-gonic/gin"
)

func RegistPath(r *gin.Engine) {
	// s2s send
	r.GET("/users/:name", controller.Users)
	// r.GET("/users/:name/outbox", controller.UsersOutbox)
	// r.GET("/users/:name/followers", controller.UsersFollowers)
	// r.GET("/users/:name/following", controller.UsersFollowing)
	// r.GET("/users/:name/collections/featured", controller.UsersCollectionsFeatured)
	// r.GET("/users/:name/collections/tags", controller.UsersCollectionsTags)
	// // https://mstdn.jp/users/meromero/statuses/110734394957749061
	// r.GET("/users/:name/statuses/:id", controller.UsersStatuses)
	// // s2s recv
	r.POST("/users/:name/inbox", controller.UsersInbox) // inbox, only on POST?
	r.POST("/inbox", controller.Inbox)                  // inbox
}

var ()
