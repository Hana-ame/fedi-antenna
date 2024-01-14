package statuses

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/Hana-ame/fedi-antenna/mastodon/model"
	"github.com/gin-gonic/gin"
)

// POST /api/v1/statuses HTTP/1.1
func Post_a_new_status(c *gin.Context) {
	var o *model.Post_a_new_status
	c.Bind(&o)

	// todo
	// should get real id
	authorizationID := utils.ParseActivitypubID("test5", "fedi.moonchan.xyz")

	err := handler.Note(authorizationID, o)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// return
}
