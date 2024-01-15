package controller

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/mastodon/controller/accounts/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/gin-gonic/gin"
)

// POST /api/v1/accounts/:id/follow HTTP/1.1
func Follow_account(c *gin.Context) {
	// REQUIRED String. The ID of the Account in the database.
	id := c.Param("id")
	// REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method.
	Authorization := c.GetHeader("Authorization")
	var data *model.Follow_account
	c.Bind(&data)
	o, err := handler.Follow_account(
		id,
		Authorization,
		data,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, o)
	return
}
