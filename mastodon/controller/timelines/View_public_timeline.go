package controller

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/gin-gonic/gin"
)

// GET /api/v1/timelines/public HTTP/1.1
func View_public_timeline(c *gin.Context) {
	// Provide this header with Bearer <user token> to gain authorized access to this API method.
	Authorization := c.GetHeader("Authorization")
	// Boolean. Show only local statuses? Defaults to false.
	local := c.Query("local")
	// Boolean. Show only remote statuses? Defaults to false.
	remote := c.Query("remote")
	// Boolean. Show only statuses with media attached? Defaults to false.
	only_media := c.Query("only_media")
	// String. All results returned will be lesser than this ID. In effect, sets an upper bound on results.
	max_id := c.Query("max_id")
	// String. All results returned will be greater than this ID. In effect, sets a lower bound on results.
	since_id := c.Query("since_id")
	// String. Returns results immediately newer than this ID. In effect, sets a cursor at this ID and paginates forward.
	min_id := c.Query("min_id")
	// Integer. Maximum number of results to return. Defaults to 20 statuses. Max 40 statuses.
	limit := c.Query("limit")
	o, err := handler.View_public_timeline(
		Authorization,
		local,
		remote,
		only_media,
		max_id,
		since_id,
		min_id,
		limit,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, o)
	return
}
