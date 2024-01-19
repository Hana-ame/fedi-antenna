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
	// lu := &core.LocalUser{
	// 	AccountID: utils.ParseActivitypubID(name, host),
	// }
	// if err := dao.Read(lu); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	o := convert.ToActivityPubUser(utils.ParseActivitypubID(name, host))

	c.JSON(http.StatusOK, o)
}
