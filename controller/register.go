package controller

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/antenna/handler"
	"github.com/Hana-ame/fedi-antenna/antenna/model"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var o *model.Register
	c.Bind(&o)

	if err := handler.Register(o); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusCreated, "")
}
