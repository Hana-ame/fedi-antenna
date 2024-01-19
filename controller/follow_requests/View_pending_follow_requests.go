package controller

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/gin-gonic/gin"
)

// GET /api/v1/follow_requests HTTP/1.1
func View_pending_follow_requests(c *gin.Context) {
	// REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method.
	Authorization := c.GetHeader("Authorization")
	// Internal parameter. Use HTTP Link header for pagination.
	max_id := c.Query("max_id")
	// Internal parameter. Use HTTP Link header for pagination.
	since_id := c.Query("since_id")
	// Integer. Maximum number of results to return. Defaults to 40 accounts. Max 80 accounts.
	limit := c.Query("limit")
	o, err := handler.View_pending_follow_requests(
		Authorization,
		max_id,
		since_id,
		limit,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, o)
	return
}
