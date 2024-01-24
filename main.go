package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/Hana-ame/fedi-antenna/Tools/myfetch"
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/controller"
	antenna "github.com/Hana-ame/fedi-antenna/antenna/controller"
	_ "github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	mastodon "github.com/Hana-ame/fedi-antenna/mastodon/controller"
	webfinger "github.com/Hana-ame/fedi-antenna/webfinger/controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Init()

	r := gin.Default()

	webfinger.RegisterPath(r)
	activitypub.RegisterPath(r)
	antenna.RegisterPath(r)
	mastodon.RegisterPath(r)

	// r.GET("/users/test7/statuses/1706100851713979", func(c *gin.Context) {
	// 	log.Println("!!!!!!")
	// 	var o map[any]any
	// 	j := `{
	// 		"@context": [
	// 				"https://www.w3.org/ns/activitystreams",
	// 				{
	// 						"ostatus": "http://ostatus.org#",
	// 						"atomUri": "ostatus:atomUri",
	// 						"inReplyToAtomUri": "ostatus:inReplyToAtomUri",
	// 						"conversation": "ostatus:conversation",
	// 						"sensitive": "as:sensitive",
	// 						"toot": "http://joinmastodon.org/ns#",
	// 						"votersCount": "toot:votersCount"
	// 				}
	// 		],
	// 		"id": "https://fedi.moonchan.xyz/users/test7/statuses/1706100851713979",
	// 		"type": "Note",
	// 		"summary": null,
	// 		"inReplyTo": null,
	// 		"published": "2024-01-24T12:54:11Z",
	// 		"url": "https://fedi.moonchan.xyz/@test7/1706100851713979",
	// 		"attributedTo": "https://fedi.moonchan.xyz/users/test7",
	// 		"to": [
	// 				"https://www.w3.org/ns/activitystreams#Public"
	// 		],
	// 		"cc": [
	// 				"https://fedi.moonchan.xyz/users/test7/followers"
	// 		],
	// 		"sensitive": false,
	// 		"atomUri": "https://fedi.moonchan.xyz/users/test7/statuses/1706100851713979",
	// 		"inReplyToAtomUri": null,
	// 		"conversation": "tag:fedi.moonchan.xyz,2024-01-24:objectId=1706100851713979:objectType=Conversation",
	// 		"content": "疯掉123",
	// 		"contentMap": {
	// 				"zh": "疯掉123"
	// 		},
	// 		"attachment": [],
	// 		"tag": [],
	// 		"replies": {
	// 				"id": "https://fedi.moonchan.xyz/users/test7/statuses/1706100851713979/replies",
	// 				"type": "Collection",
	// 				"first": {
	// 						"type": "CollectionPage",
	// 						"next": "https://fedi.moonchan.xyz/users/test7/statuses/1706100851713979/replies?only_other_accounts=true&page=true",
	// 						"partOf": "https://fedi.moonchan.xyz/users/test7/statuses/1706100851713979/replies",
	// 						"items": []
	// 				}
	// 		}
	// }`
	// 	json.Unmarshal([]byte(j), &o)
	// 	c.JSON(http.StatusOK, o)
	// })

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
