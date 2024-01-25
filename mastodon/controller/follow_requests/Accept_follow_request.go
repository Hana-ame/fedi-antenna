package controller

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/gin-gonic/gin"
)

// POST /api/v1/follow_requests/:account_id/authorize HTTP/1.1
func Accept_follow_request(c *gin.Context) {
	// REQUIRED String. The ID of the Account in the database.
	account_id := c.Param("account_id")
	// REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method.
	Authorization := c.GetHeader("Authorization")
	o, err := handler.Accept_follow_request(
		account_id,
		Authorization,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, o)
	return
}
