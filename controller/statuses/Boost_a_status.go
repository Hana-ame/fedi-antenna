package statuses

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/Hana-ame/fedi-antenna/mastodon/model"
	"github.com/gin-gonic/gin"
)

// POST /api/v1/statuses/:id/reblog HTTP/1.1
func Boost_a_status(c *gin.Context) {
	var o *model.Post_a_new_status
	c.Bind(&o)

	authorizationID := utils.ParseActivitypubID("test5", "fedi.moonchan.xyz")
	id := c.Param("id")

	status, err := handler.Reblog(authorizationID, id, o.Visibility)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, status)
	return

}
