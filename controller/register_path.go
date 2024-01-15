package mastodon

import (
	accounts "github.com/Hana-ame/fedi-antenna/mastodon/controller/accounts"
	statuses "github.com/Hana-ame/fedi-antenna/mastodon/controller/statuses"
	"github.com/gin-gonic/gin"
)

func RegisterPath(r *gin.Engine) {
	// account
	r.POST("/api/v1/accounts/:id/follow", accounts.Follow_account)
	r.POST("/api/v1/accounts/:id/unfollow", accounts.Unfollow_account)
	// statuses
	r.POST("/api/v1/statuses", statuses.Post_a_new_status)
	r.DELETE("/api/v1/statuses/:id", statuses.Delete_a_status)
	// reblog
	r.POST("/api/v1/statuses/:id/reblog", statuses.Boost_a_status)
	r.POST("/api/v1/statuses/:id/unreblog", statuses.Undo_boost_of_a_status)
	// favourite
	r.POST("/api/v1/statuses/:id/favourite", statuses.Favourite_a_status)
	r.POST("/api/v1/statuses/:id/unfavourite", statuses.Undo_favourite_of_a_status)
}
