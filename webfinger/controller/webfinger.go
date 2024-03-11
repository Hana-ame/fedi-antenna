package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"

	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/webfinger/model"
)

func Webfinger(c *gin.Context) {
	// 被访问
	// 取得query
	resource := c.Query("resource")
	if !strings.HasPrefix(resource, "acct:") {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "format error"})
		return
	}
	// 取出成为 uesrname 和 host
	acct := strings.TrimPrefix(resource, "acct:")
	username, host := utils.ParseUserAndHost(acct)
	// 判断是否存在 not used
	user := &core.LocalUser{
		ActivitypubID: utils.NameAndHost2ActivitypubID(username, host),
	}
	if err := dao.Read(dao.DB(), user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		_ = user

		// 形成返回数据
		// todo: alias.
		webfingerObj, err := model.NewWebfingerObj(username, host)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, webfingerObj)
	}
}
