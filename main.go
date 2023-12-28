package main

import (
	"github.com/Hana-ame/fedi-antenna/activitypub"
	_ "github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/webfinger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	godotenv.Load(".env")

	r := gin.Default()

	webfinger.RegistPath(r)
	activitypub.RegistPath(r)

	r.Run("0.0.0.0:3000")
}
