package main

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/actions"
	"github.com/Hana-ame/fedi-antenna/controller"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.GET("/webfinger", controller.Webfinger)
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	actions.FetchWebfingerByAcct("")

	r.Run()
}
