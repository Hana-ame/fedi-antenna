package controller

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"github.com/gin-gonic/gin"
)

// /users/:name/statuses/:id
func UsersStatus(c *gin.Context) {
	// name := c.Param("name")
	// host := c.Request.Host
	id := c.Param("id")

	status := &entities.Status{
		Id: id,
	}
	if err := dao.Read(dao.DB(), status); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	o := convert.StatusToNote(status)
	c.Writer.Header().Set("Content-Type", "application/activity+json; charset=utf-8")
	c.JSON(http.StatusOK, o)
}
