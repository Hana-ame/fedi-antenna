package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Webfinger(c *gin.Context) {
	c.JSONP(http.StatusOK, map[string]string{"code": http.ErrNotSupported.ErrorString})
}
