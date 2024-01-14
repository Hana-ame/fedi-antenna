package mastodon

import (
	"github.com/Hana-ame/fedi-antenna/mastodon/controller/statuses"
	"github.com/gin-gonic/gin"
)

func RegisterPath(r *gin.Engine) {
	// statuses
	r.POST("/api/v1/statuses", statuses.Post_a_new_status) // inbox
}
