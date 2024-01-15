package statuses

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/mastodon/controller/statuses/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/gin-gonic/gin"
)

// POST /api/v1/statuses/:id/reblog HTTP/1.1
func Boost_a_status(c *gin.Context) {
	// REQUIRED String. The ID of the Status in the database.
	id := c.Param("id")
	// REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method.
	Authorization := c.GetHeader("Authorization")
	var data *model.Boost_a_status
	c.Bind(&data)
	o, err := handler.Boost_a_status(
		id,
		Authorization,
		data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, o)
	return
}
