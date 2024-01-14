package statuses

import (
	"net/http"

	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
	"github.com/gin-gonic/gin"
)

// DELETE /api/v1/statuses/:id HTTP/1.1
func Delete_a_status(c *gin.Context) {
	authorizationID := utils.ParseActivitypubID("test5", "fedi.moonchan.xyz")
	id := c.Param("id")

	err := handler.DeleteNote(authorizationID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
}
