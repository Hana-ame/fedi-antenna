package controller

import (
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/Hana-ame/fedi-antenna/mastodon/model"
	"github.com/gin-gonic/gin"
)

// POST /api/v1/statuses HTTP/1.1
func Post_a_new_status(c *gin.Context) {
	var o *model.Status
	c.Bind(&o)

	// todo
	// should get real id
	id := utils.ParseActivitypubID("test5", "fedi.moonchan.xyz")

	handler.Note(id, o)
}
