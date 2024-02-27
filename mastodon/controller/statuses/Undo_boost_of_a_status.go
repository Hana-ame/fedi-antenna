package statuses

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/gin-gonic/gin"
)

// POST /api/v1/statuses/:id/unreblog HTTP/1.1
func Undo_boost_of_a_status(c *gin.Context) {
	// REQUIRED String. The ID of the Status in the database.
	id := c.Param("id")
	// REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method.
	Authorization := c.GetHeader("Authorization")
	o, err := handler.Undo_boost_of_a_status(
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
