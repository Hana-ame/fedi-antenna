package main

import (
	"net/http"
	"strings"

	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/utils"
	"github.com/gin-gonic/gin"
)

// webfinger 的 server
func webfinger(c *gin.Context) {
	// 被访问
	// 取得query
	resource := c.Query("resource")
	if !strings.HasPrefix(resource, "acct:") {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "format error"})
		return
	}
	// 取出成为 uesrname 和 host
	acct := strings.TrimPrefix(resource, "acct:")
	username, host := utils.ParseAcctStrToUserAndHost(acct)
	// 判断是否存在 not used
	// if !core.IsUserExist(username, host) {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	// 	return
	// }

	// 形成返回数据
	webfingerObj, err := core.WebfingerObj(username, host)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, webfingerObj)
}
