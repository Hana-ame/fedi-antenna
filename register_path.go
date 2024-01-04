package antenna

import (
	"github.com/Hana-ame/fedi-antenna/antenna/controller"
	"github.com/gin-gonic/gin"
)

func RegisterPath(r *gin.Engine) {
	r.POST("/register", controller.Register) // inbox
}
