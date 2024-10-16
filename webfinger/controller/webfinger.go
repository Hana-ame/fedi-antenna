package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
	"github.com/Hana-ame/fedi-antenna/activitypub/dao"
	"github.com/Hana-ame/fedi-antenna/activitypub/dao/model"
	"github.com/Hana-ame/fedi-antenna/utils"
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
	user := &model.User{
		ID: utils.ParseActivitypubID(username, host),
	}
	if err := dao.Read(user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		// 形成返回数据
		// todo: alias.
		webfingerObj, err := CreateWebfingerObj(username, host)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, webfingerObj)
	}
}

func CreateWebfingerObj(username, host string) (o *orderedmap.OrderedMap, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%s", e)
		}
	}()
	o = orderedmap.New()
	o.Set("subject", "acct:"+username+"@"+host)
	o.Set("links", []*orderedmap.OrderedMap{
		utils.NewMapFromKV([]*utils.KV{
			{Key: "rel", Value: "self"},
			{Key: "type", Value: "application/activity+json"},
			{Key: "href", Value: "https://" + host + "/users/" + username},
		}),
		utils.NewMapFromKV([]*utils.KV{
			{Key: "rel", Value: "http://webfinger.net/rel/profile-page"},
			{Key: "type", Value: "text/html"},
			{Key: "href", Value: "https://" + host + "/@" + username},
		}),
		utils.NewMapFromKV([]*utils.KV{ // dunno what it is
			{Key: "rel", Value: "http://ostatus.org/schema/1.0/subscribe"},
			{Key: "template", Value: "https://p1.a9z.dev/authorize-follow?acct={uri}"},
		}),
	})
	return
}
