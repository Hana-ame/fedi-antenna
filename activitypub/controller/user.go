package controller

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/gin-gonic/gin"
)

// /users/:name
func Users(c *gin.Context) {
	name := c.Param("name")
	host := c.Request.Host

	o := convert.ToActivityPubUser(utils.NameAndHost2ActivitypubID(name, host))

	c.Writer.Header().Set("Content-Type", "application/activity+json; charset=utf-8")
	c.JSON(http.StatusOK, o)
}
