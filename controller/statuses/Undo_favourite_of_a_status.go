package statuses

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/gin-gonic/gin"
)

// POST /api/v1/statuses/:id/unfavourite HTTP/1.1
func Undo_favourite_of_a_status(c *gin.Context) {
	authorizationID := utils.ParseActivitypubID("test5", "fedi.moonchan.xyz")
	id := c.Param("id")

	status, err := handler.Unfavourite(authorizationID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, status)
	return
}
