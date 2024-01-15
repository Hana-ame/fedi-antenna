package statuses

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/mastodon/controller/statuses/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/gin-gonic/gin"
)

// POST /api/v1/statuses HTTP/1.1
func Post_a_new_status(c *gin.Context) {
	// REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method.
	Authorization := c.GetHeader("Authorization")
	// Provide this header with any arbitrary string to prevent duplicate submissions of the same status. Consider using a hash or UUID generated client-side. Idempotency keys are stored for up to 1 hour.
	IdempotencyKey := c.GetHeader("Idempotency-Key")

	var data *model.Post_a_new_status
	c.Bind(&data)

	status, err := handler.Post_a_new_status(Authorization, IdempotencyKey, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, status)
	return

}
