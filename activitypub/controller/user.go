package controller

import (
	"log"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"github.com/gin-gonic/gin"
)

// /users/:name
func Users(c *gin.Context) {
	name := c.Param("name")
	host := c.Request.Host

	activitypubID := utils.NameAndHost2ActivitypubID(name, host)
	localuser := &model.LocalUser{
		ActivitypubID: activitypubID,
	}
	if err := dao.Read(dao.DB(), localuser); err != nil {
		log.Printf("%s", err.Error())
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	acct := &entities.Account{
		Uri: activitypubID,
	}
	if err := dao.Read(dao.DB(), acct); err != nil {
		log.Printf("%s", err.Error())
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	o := convert.AccountToActivitypubUser(acct, localuser)

	c.Writer.Header().Set("Content-Type", "application/activity+json; charset=utf-8")
	c.JSON(http.StatusOK, o)
}
