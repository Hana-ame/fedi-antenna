package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/Hana-ame/fedi-antenna/Tools/myfetch"
	"github.com/Hana-ame/fedi-antenna/activitypub"
	"github.com/Hana-ame/fedi-antenna/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Init()

	r := gin.Default()

	activitypub.RegisterPath(r)

	r.Run("0.0.0.0:3000")
}

// that will work in debug mode.
func init() {
	Init()
}

func Init() {
	// set log
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("is that ok?")
	// set proxy
	proxyUrl, _ := url.Parse("http://DESKTOP-LLULJ2Q.mshome.net:10809")
	myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	myfetch.SetClients([]*http.Client{myClient})

	godotenv.Load(".env")

	// start up
	utils.AliasMap["localhost:3000"] = "fedi.moonchan.xyz"

}
