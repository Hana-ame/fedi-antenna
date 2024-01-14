package statuses

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/gin-gonic/gin"
)

// todo: the body do have some parameters, not implemented.
// POST /api/v1/accounts/:id/follow HTTP/1.1
func Follow_account(c *gin.Context) {
	authorizationID := utils.ParseActivitypubID("test5", "fedi.moonchan.xyz")
	id := c.Param("id")

	relationship, err := handler.Follow(authorizationID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, relationship)
}
