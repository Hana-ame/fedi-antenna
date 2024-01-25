package controller

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"github.com/gin-gonic/gin"
)

// /users/:name
func Users(c *gin.Context) {
	name := c.Param("name")
	host := c.Request.Host

	o := convert.ToActivityPubUser(utils.ParseActivitypubID(name, host))

	c.JSON(http.StatusOK, o)
}

func UsersStatuses(c *gin.Context) {
	// name := c.Param("name")
	// host := c.Request.Host
	id := c.Param("id")

	status := &entities.Status{}
	if tx := dao.Where("Id = ?", id).First(status); tx.Error != nil {
		c.JSON(http.StatusInternalServerError, tx.Error)
	}

	o := convert.ToActivityPubNote(status)
	c.Writer.Header().Set("Content-Type", "application/activity+json; charset=utf-8")
	c.JSON(http.StatusOK, o)
}
