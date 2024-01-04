package controller

import (
	"net/http"

	handler "github.com/Hana-ame/fedi-antenna/core/handler/antenna"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var username string
	var host string
	var email string
	var passwd string

	p := &utils.Parser{Context: c}

	if err := p.PostForm("username", &username); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.PostForm("email", &email); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.PostForm("host", &host); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := p.PostForm("passwd", &passwd); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := handler.Register(
		username,
		host,
		email,
		passwd,
	); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusCreated, "")
}
