package activitypub

import (
	"github.com/Hana-ame/fedi-antenna/utils"
	"github.com/iancoleman/orderedmap"
)

var UserObjAtContext = []any{
	"https://www.w3.org/ns/activitystreams",
	"https://w3id.org/security/v1",
	utils.OrderedMap([]string{
		"manuallyApprovesFollowers",
		"toot",
		"featured",
		"featuredTags",
		"alsoKnownAs",
		"movedTo",
		"schema",
		"PropertyValue",
		"value",
		"discoverable",
		"Device",
		"Ed25519Signature",
		"Ed25519Key",
		"Curve25519Key",
		"EncryptedMessage",
		"publicKeyBase64",
		"deviceId",
		"claim",
		"fingerprintKey",
		"identityKey",
		"devices",
		"messageFranking",
		"messageType",
		"cipherText",
		"suspended",
		"focalPoint",
	}, []any{
		"as:manuallyApprovesFollowers",
		"http://joinmastodon.org/ns#",
		utils.IdType("toot:featured", "@id"),
		utils.IdType("toot:featuredTags", "@id"),
		utils.IdType("toot:alsoKnownAs", "@id"),
		utils.IdType("toot:movedTo", "@id"),
		"http://schema.org#",
		"schema:PropertyValue",
		"schema:value",
		"toot:discoverable",
		"toot:Device",
		"toot:Ed25519Signature",
		"toot:Ed25519Key",
		"toot:Curve25519Key",
		"toot:EncryptedMessage",
		"toot:publicKeyBase64",
		"toot:deviceId",
		utils.IdType("toot:claim", "@id"),
		utils.IdType("toot:fingerprintKey", "@id"),
		utils.IdType("toot:identityKey", "@id"),
		utils.IdType("toot:devices", "@id"),
		"toot:messageFranking",
		"toot:messageType",
		"toot:cipherText",
		"toot:suspended",
		utils.OrderedMap([]string{"@id", "@container"}, []any{"toot:focalPoint", "@list"}),
	}),
}

func UserObj(
	host, name string,
	published int64, // timestamp in us,
	pubkey, icon *orderedmap.OrderedMap,
) *orderedmap.OrderedMap {
	o := utils.OrderedMap(
		[]string{
			"@context",
			"id",
			"type",
			"following",
			"followers",
			"inbox",
			"outbox",
			"featured",
			"featuredTags",
			"preferredUsername",
			"name",
			"summary",
			"url",
			"manuallyApprovesFollowers",
			"discoverable",
			"published",
			"devices",
			"publicKey",
			"tag",
			"attachment",
			"endpoints",
			"icon",
		}, []any{
			UserObjAtContext,
			APUserID(name, host),
			"Person",
			APUserID(name, host) + "/following",
			APUserID(name, host) + "/followers",
			APUserID(name, host) + "/inbox",
			APUserID(name, host) + "/outbox",
			APUserID(name, host) + "/collections/featured",
			APUserID(name, host) + "/collections/tags",
			name,
			"",
			"",
			"https://" + host + "/@" + name,
			false,
			false,
			utils.TimestampToRFC3339(published),
			APUserID(name, host) + "/collections/devices",
			pubkey,                     // public key
			[]string{},                 // tag
			[]*orderedmap.OrderedMap{}, // attachment (tags)
			utils.OrderedMap([]string{"sharedInbox"}, []any{"https://" + host + "/inbox"}), // endpoints
			icon, // icon
		})
	return o
}

// "https://" + host + "/users/" + user
func APUserID(user, host string) string {
	return "https://" + host + "/users/" + user
}
