package main

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/webfinger"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	webfinger.RegistPath(r)

	id, err := webfinger.GetUserIdFromAcct("meromero@p1.a9z.dev")
	log.Println(id)
	log.Println(err)

	r.Run("0.0.0.0:8080")
}
