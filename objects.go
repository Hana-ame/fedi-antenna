package main

import (
	tools "github.com/Hana-ame/fedi-antenna/Tools"
	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
)

// utils
func newOrderedMap(keys []string, values []interface{}) *orderedmap.OrderedMap {
	o := orderedmap.New()
	for i, key := range keys {
		o.Set(key, values[i])
	}
	return o
}

// following
func genFollowObj(actor, object string) *orderedmap.OrderedMap {
	o := newOrderedMap(
		[]string{
			"@context",
			"id",
			"type",
			"actor",
			"object",
		}, []any{
			"https://www.w3.org/ns/activitystreams",
			"https://" + host + "/something",
			"Follow",
			actor,
			object,
		})
	return o
}

// user
func genUserObj(name string) *orderedmap.OrderedMap {
	o := newOrderedMap(
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
			genASContextObj(),
			"https://" + host + "/users/" + name,
			"Person",
			"https://" + host + "/users/" + name + "/following",
			"https://" + host + "/users/" + name + "/followers",
			"https://" + host + "/users/" + name + "/inbox",
			"https://" + host + "/users/" + name + "/outbox",
			"https://" + host + "/users/" + name + "/collections/featured",
			"https://" + host + "/users/" + name + "/collections/tags",
			name,
			"",
			"",
			"https://" + host + "/@" + name,
			false,
			false,
			"2023-01-01T00:00:00Z",
			"https://" + host + "/users/" + name + "/collections/devices",
			newOrderedMap(
				[]string{
					"id",
					"owner",
					"publicKeyPem",
				}, []any{
					"https://" + host + "/users/" + name + "#main-key",
					"https://" + host + "/users/" + name,
					genPkPem(name),
				}), // public key
			[]string{},                 // tag
			[]*orderedmap.OrderedMap{}, // attachment (tags)
			newOrderedMap([]string{"sharedInbox"}, []any{"https://" + host + "/inbox"}), // endpoints
			newOrderedMap([]string{
				"type",
				"mediaType",
				"url",
			}, []any{
				"Image",
				"image/jpeg",
				"https://s3.arkjp.net/misskey/678ad158-f160-48f4-a369-8756aa92350e.jpg",
			}), // icon
		})
	return o
}

func genPkPem(name string) string {
	pk, _ := tools.ReadKeyFromFile(name + ".pem")
	pubKeyStr, _ := tools.MarshalPublicKey(&pk.PublicKey)
	return string(pubKeyStr)
}

func genASContextObj() []any {
	return []any{
		"https://www.w3.org/ns/activitystreams",
		"https://w3id.org/security/v1",
		newOrderedMap([]string{
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
			newOrderedMap([]string{"@id", "@type"}, []any{"toot:featured", "@id"}),
			newOrderedMap([]string{"@id", "@type"}, []any{"toot:featuredTags", "@id"}),
			newOrderedMap([]string{"@id", "@type"}, []any{"toot:alsoKnownAs", "@id"}),
			newOrderedMap([]string{"@id", "@type"}, []any{"toot:movedTo", "@id"}),
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
			newOrderedMap([]string{"@id", "@type"}, []any{"toot:claim", "@id"}),
			newOrderedMap([]string{"@id", "@type"}, []any{"toot:fingerprintKey", "@id"}),
			newOrderedMap([]string{"@id", "@type"}, []any{"toot:identityKey", "@id"}),
			newOrderedMap([]string{"@id", "@type"}, []any{"toot:devices", "@id"}),
			"toot:messageFranking",
			"toot:messageType",
			"toot:cipherText",
			"toot:suspended",
			newOrderedMap([]string{"@id", "@container"}, []any{"toot:focalPoint", "@list"}),
		}),
	}
}
