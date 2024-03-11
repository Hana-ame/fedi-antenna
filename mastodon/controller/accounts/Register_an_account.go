package controller

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/mastodon/controller/accounts/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/gin-gonic/gin"
)

// POST /api/v1/accounts HTTP/1.1
func Register_an_account(c *gin.Context) {
	// REQUIRED Provide this header with Bearer <app token> to gain authorized access to this API method.
	Authorization := c.GetHeader("Authorization")
	host := c.GetHeader("Host")
	var data *model.Register_an_account
	c.Bind(&data)
	err := handler.Register_an_account(
		Authorization,
		host,
		data,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
	return
}
