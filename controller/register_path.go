package controller

import (
	"github.com/gin-gonic/gin"
)

func RegisterPath(r *gin.Engine) {
	r.GET("/.well-known/webfinger", Webfinger)
}
