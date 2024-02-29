package model

import "github.com/Hana-ame/fedi-antenna/core/utils"

type User struct {
	Context any `json:"@context,omitempty" gorm:"-"`

	// it is a url
	ID string `json:"id" gorm:"primarykey"`
	// helper.
	// Account string `json:"-"`

	// "Person"
	Type string `json:"type" gorm:"-"`

	Following    string `json:"following"`
	Followers    string `json:"followers"`
	Inbox        string `json:"inbox"`
	Outbox       string `json:"outbox"`
	Featured     string `json:"featured"`
	FeaturedTags string `json:"featuredTags"`

	// without @host
	PreferredUsername string `json:"preferredUsername" gorm:"type:text collate nocase"`

	// could be empty
	// the name which shows on profile
	Name string `json:"name"`

	Summary string `json:"summary"`

	// homepage/profilepage that access by browser
	URL string `json:"url"`

	ManuallyApprovesFollowers bool `json:"manuallyApprovesFollowers"`
	Discoverable              bool `json:"discoverable"`

	// TimestampToRFC3339
	Published string `json:"published"`

	// also a url
	Devices     string     `json:"devices"`
	AlsoKnownAs []string   `json:"alsoKnownAs,omitempty" gorm:"serializer:json"`
	PublicKey   *PublicKey `json:"publicKey" gorm:"foreignKey:Owner;references:ID"`

	// the icon / images in :name:
	Tag []*Tag `json:"tag" gorm:"-"` // todo

	// what is the type?
	Attachment []any `json:"attachment" gorm:"-"` // todo

	SharedInbox string            `json:"sharedInbox,omitempty"`
	Endpoint    map[string]string `json:"endpoints" gorm:"-"`

	// the url of Image
	IconURL string `json:"-"`
	Icon    *Image `json:"icon" gorm:"foreignKey:IconURL;references:URL"`
}

var UserContext = []any{
	"https://www.w3.org/ns/activitystreams",
	"https://w3id.org/security/v1",
	utils.NewMapFromKV([]*utils.KV{
		{Key: "manuallyApprovesFollowers", Value: "as:manuallyApprovesFollowers"},
		{Key: "toot", Value: "http://joinmastodon.org/ns#"},
		{Key: "featured", Value: IDType{"toot:featured", "@id"}},
		{Key: "featuredTags", Value: IDType{"toot:featuredTags", "@id"}},
		{Key: "alsoKnownAs", Value: IDType{"toot:alsoKnownAs", "@id"}},
		{Key: "movedTo", Value: IDType{"toot:movedTo", "@id"}},
		{Key: "schema", Value: "http://schema.org#"},
		{Key: "PropertyValue", Value: "schema:PropertyValue"},
		{Key: "value", Value: "schema:value"},
		{Key: "discoverable", Value: "toot:discoverable"},
		{Key: "Device", Value: "toot:Device"},
		{Key: "Ed25519Signature", Value: "toot:Ed25519Signature"},
		{Key: "Ed25519Key", Value: "toot:Ed25519Key"},
		{Key: "Curve25519Key", Value: "toot:Curve25519Key"},
		{Key: "EncryptedMessage", Value: "toot:EncryptedMessage"},
		{Key: "publicKeyBase64", Value: "toot:publicKeyBase64"},
		{Key: "deviceId", Value: "toot:deviceId"},
		{Key: "claim", Value: IDType{"toot:claim", "@id"}},
		{Key: "fingerprintKey", Value: IDType{"toot:fingerprintKey", "@id"}},
		{Key: "identityKey", Value: IDType{"toot:identityKey", "@id"}},
		{Key: "devices", Value: IDType{"toot:devices", "@id"}},
		{Key: "messageFranking", Value: "toot:messageFranking"},
		{Key: "messageType", Value: "toot:messageType"},
		{Key: "cipherText", Value: "toot:cipherText"},
		{Key: "suspended", Value: "toot:suspended"},
		{Key: "focalPoint", Value: map[string]string{"@id": "toot:focalPoint", "@container": "@list"}},
	}),
}

func NewUser(name, host string) *User {
	return &User{
		Context:           UserContext,
		ID:                utils.NameAndHost2ProfileUrlActivitypubID(name, host),
		Type:              "Person",
		Following:         utils.NameAndHost2ProfileUrlActivitypubID(name, host) + "/following",
		Followers:         utils.NameAndHost2ProfileUrlActivitypubID(name, host) + "/followers",
		Inbox:             utils.NameAndHost2ProfileUrlActivitypubID(name, host) + "/inbox",
		Outbox:            utils.NameAndHost2ProfileUrlActivitypubID(name, host) + "/outbox",
		Featured:          utils.NameAndHost2ProfileUrlActivitypubID(name, host) + "/collections/featured",
		FeaturedTags:      utils.NameAndHost2ProfileUrlActivitypubID(name, host) + "/collections/tags",
		PreferredUsername: name,

		URL: utils.NameAndHost2ProfileUrl(name, host),

		Published: utils.TimestampToRFC3339(utils.NewTimestamp()),
		Devices:   utils.NameAndHost2ProfileUrlActivitypubID(name, host) + "/collections/devices",
		PublicKey: NewPublicKey(utils.NameAndHost2ProfileUrlActivitypubID(name, host)),
		Endpoint:  map[string]string{"sharedInbox": "https://" + host + "/inbox"},

		Icon: &Image{
			"Image",
			"image/png",
			"https://twimg.moonchan.xyz/media/GB8hT5vacAAK6wa?format=png&name=medium",
		},
	}
}

func (o *User) Autofill() {
	// get name and host
	name, host := utils.ActivitypubID2NameAndHost(o.ID)

	o.Context = UserContext
	if o.Type == "" {
		o.Type = "Person"
	}
	if o.Following == "" {
		o.Following = o.ID + "/following"
	}
	if o.Followers == "" {
		o.Followers = o.ID + "/followers"
	}
	if o.Inbox == "" {
		o.Inbox = o.ID + "/inbox"
	}
	if o.Outbox == "" {
		o.Outbox = o.ID + "/outbox"
	}
	if o.Featured == "" {
		o.Featured = o.ID + "/collections/featured"
	}
	if o.FeaturedTags == "" {
		o.FeaturedTags = o.ID + "/collections/tags"
	}
	if o.PreferredUsername == "" {
		o.PreferredUsername = name
	}
	if o.URL == "" {
		o.URL = utils.NameAndHost2ProfileUrl(name, host)
	}
	if o.Devices == "" {
		o.Devices = o.ID + "/collections/devices"
	}
	if o.SharedInbox == "" {
		o.SharedInbox = "https://" + host + "/inbox"
	}
	if o.Endpoint == nil {
		o.Endpoint = map[string]string{"sharedInbox": "https://" + host + "/inbox"}
	}
}

type IDType struct {
	ID   string `json:"@id"`
	Type string `json:"@type"`
}
