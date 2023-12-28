package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/utils"

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
	if !core.IsAccountExist(username, host) {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	// 形成返回数据
	webfingerObj, err := model.CreateWebfingerObj(username, host)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, webfingerObj)
}
