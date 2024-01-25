package controller

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/Hana-ame/orderedmap"
)

func TestXxx(t *testing.T) {
	b := []byte(`{"@context":["https://www.w3.org/ns/activitystreams",{"ostatus":"http://ostatus.org#","atomUri":"ostatus:atomUri","inReplyToAtomUri":"ostatus:inReplyToAtomUri","conversation":"ostatus:conversation","sensitive":"as:sensitive","toot":"http://joinmastodon.org/ns#","votersCount":"toot:votersCount"}],"id":"https://me.ns.ci/users/cocoon/statuses/111707195784478139/activity","type":"Create","actor":"https://me.ns.ci/users/cocoon","published":"2024-01-06T04:48:57Z","to":["https://fedi.moonchan.xyz/users/test4"],"cc":[],"object":{"id":"https://me.ns.ci/users/cocoon/statuses/111707195784478139","type":"Note","summary":null,"inReplyTo":null,"published":"2024-01-06T04:48:57Z","url":"https://me.ns.ci/@cocoon/111707195784478139","attributedTo":"https://me.ns.ci/users/cocoon","to":["https://fedi.moonchan.xyz/users/test4"],"cc":[],"sensitive":false,"atomUri":"https://me.ns.ci/users/cocoon/statuses/111707195784478139","inReplyToAtomUri":null,"conversation":"tag:me.ns.ci,2024-01-06:objectId=39533213:objectType=Conversation","content":"\u003cp\u003e\u003cspan class=\"h-card\" translate=\"no\"\u003e\u003ca href=\"https://fedi.moonchan.xyz/@test4\" class=\"u-url mention\"\u003e@\u003cspan\u003etest4\u003c/span\u003e\u003c/a\u003e\u003c/span\u003e 123\u003c/p\u003e","contentMap":{"zh":"\u003cp\u003e\u003cspan class=\"h-card\" translate=\"no\"\u003e\u003ca href=\"https://fedi.moonchan.xyz/@test4\" class=\"u-url mention\"\u003e@\u003cspan\u003etest4\u003c/span\u003e\u003c/a\u003e\u003c/span\u003e 123\u003c/p\u003e"},"attachment":[],"tag":[{"type":"Mention","href":"https://fedi.moonchan.xyz/users/test4","name":"@test4@fedi.moonchan.xyz"}],"replies":{"id":"https://me.ns.ci/users/cocoon/statuses/111707195784478139/replies","type":"Collection","first":{"type":"CollectionPage","next":"https://me.ns.ci/users/cocoon/statuses/111707195784478139/replies?only_other_accounts=true\u0026page=true","partOf":"https://me.ns.ci/users/cocoon/statuses/111707195784478139/replies","items":[]}}}}`)
	o := orderedmap.New()
	json.Unmarshal(b, &o)
	log.Printf("%+v", o)
	io := o.GetOrDefault("object", orderedmap.New())
	oo, ok := io.(orderedmap.OrderedMap)
	log.Printf("%+v, %v", oo, ok)

}
