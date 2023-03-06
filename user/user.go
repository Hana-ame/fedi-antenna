package user

import (
	"fmt"

	"github.com/Hana-ame/fedi-antenna/httpsig"
	"github.com/ake-persson/mapslice-json"
)

var domain = "v.meromeromeiro.top"

func UserAS(username string) ([]byte, error) {

	pk, err := httpsig.ReadKeyFromFile("privateKey.pem")
	if err != nil {
		return nil, err
	}

	ms := mapslice.MapSlice{
		mapslice.MapItem{Key: "@context", Value: []any{
			"https://www.w3.org/ns/activitystreams",
			"https://w3id.org/security/v1",
			mapslice.MapSlice{
				msKV("manuallyApprovesFollowers", "as:manuallyApprovesFollowers"),
				msKV("sensitive", "as:sensitive"),
				msKV("Hashtag", "as:Hashtag"),
				msKV("quoteUrl", "as:quoteUrl"),
				msKV("toot", "http://joinmastodon.org/ns#"),
				msKV("Emoji", "toot:Emoji"),
				msKV("featured", "toot:featured"),
				msKV("discoverable", "toot:discoverable"),
				msKV("schema", "http://schema.org#"),
				msKV("PropertyValue", "schema:PropertyValue"),
				msKV("value", "schema:value"),
				msKV("misskey", "https://misskey-hub.net/ns#"),
				msKV("_misskey_content", "misskey:_misskey_content"),
				msKV("_misskey_quote", "misskey:_misskey_quote"),
				msKV("_misskey_reaction", "misskey:_misskey_reaction"),
				msKV("_misskey_votes", "misskey:_misskey_votes"),
				msKV("isCat", "misskey:isCat"),
				msKV("vcard", "http://www.w3.org/2006/vcard/ns#"),
			},
		}},
		mapslice.MapItem{Key: "type", Value: "Person"},
		mapslice.MapItem{Key: "id", Value: fmt.Sprintf("https://%s/users/%s", domain, username)},
		mapslice.MapItem{Key: "inbox", Value: fmt.Sprintf("https://%s/users/%s/inbox", domain, username)},
		mapslice.MapItem{Key: "outbox", Value: fmt.Sprintf("https://%s/users/%s/outbox", domain, username)},
		mapslice.MapItem{Key: "followers", Value: fmt.Sprintf("https://%s/users/%s/followers", domain, username)},
		mapslice.MapItem{Key: "following", Value: fmt.Sprintf("https://%s/users/%s/following", domain, username)},
		mapslice.MapItem{Key: "featured", Value: fmt.Sprintf("https://%s/users/%s/collections/featured", domain, username)},
		mapslice.MapItem{Key: "sharedInbox", Value: fmt.Sprintf("https://%s/inbox", domain)},
		mapslice.MapItem{Key: "endpoints", Value: mapslice.MapSlice{
			mapslice.MapItem{Key: "sharedInbox", Value: fmt.Sprintf("https://%s/inbox", domain)},
		}},
		msKV("url", fmt.Sprintf("https://%s/@%s", domain, username)),
		msKV("preferredUsername", username),
		msKV("name", nil),
		msKV("summary", nil),
		msKV("icon", nil),
		msKV("image", nil),
		msKV("tag", []any{}),
		msKV("manuallyApprovesFollowers", false),
		msKV("discoverable", true),
		msKV("publicKey", mapslice.MapSlice{
			msKV("id", fmt.Sprintf("https://%s/users/%s#main-key", domain, username)),
			msKV("type", "Key"),
			msKV("owner", fmt.Sprintf("https://%s/users/%s", domain, username)),
			msKV("publicKeyPem", string(httpsig.MarshalPublicKey(&pk.PublicKey))),
		}),
	}

	return ms.MarshalJSON()
}

func msKV(key, value any) mapslice.MapItem {
	return mapslice.MapItem{Key: key, Value: value}
}
