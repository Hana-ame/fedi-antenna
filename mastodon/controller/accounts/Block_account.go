package controller

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/gin-gonic/gin"
)

// POST /api/v1/accounts/:id/block HTTP/1.1
func Block_account(c *gin.Context) {
	// REQUIRED String. The ID of the Account in the database.
	id := c.Param("id")
	// REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method.
	Authorization := c.GetHeader("Authorization")
	o, err := handler.Block_account(
		id,
		Authorization,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, o)
	return
}
