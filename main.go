package webfinger

import (
	"github.com/Hana-ame/fedi-antenna/webfinger/actions"
	"github.com/Hana-ame/fedi-antenna/webfinger/controller"
	"github.com/gin-gonic/gin"
)

func RegistPath(r *gin.Engine) {
	r.GET("/.well-known/webfinger", controller.Webfinger)
}

var (
	// methods
	FetchWebfingerByAcct = actions.FetchWebfingerByAcct
	GetUserIdFromAcct    = actions.GetUserIdFromAcct
)
