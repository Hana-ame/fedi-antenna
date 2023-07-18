package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func users(c *gin.Context) {
	name := c.Param("name")
	if name != "nanaka" {
		c.JSON(http.StatusNotFound, gin.H{"error": "only support nanaka@fedi.moonchan.xyz"})
		return
	}
	c.JSON(http.StatusOK, genUserObj(name))
}
