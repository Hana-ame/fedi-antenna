package statuses

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/gin-gonic/gin"
)

// GET /api/v1/statuses/:id HTTP/1.1
func View_a_single_status(c *gin.Context) {
	// REQUIRED String. The ID of the Status in the database.
	id := c.Param("id")
	// Provide this header with Bearer <user token> to gain authorized access to this API method.
	Authorization := c.GetHeader("Authorization")
	o, err := handler.View_a_single_status(
		id,
		Authorization,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, o)
	return
}
