package controller

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/gin-gonic/gin"
)

// /users/:name
func Users(c *gin.Context) {
	name := c.Param("name")
	host := c.Request.Host
	o, err := core.ReadActivitypubUser(name, host)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if o == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	}
	c.JSON(http.StatusOK, o)
}
