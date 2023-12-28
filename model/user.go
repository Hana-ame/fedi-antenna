package model

import "github.com/Hana-ame/fedi-antenna/core/utils"

type User struct {
	Context []any `json:"@context" gorm:"-"`

	// it is a url
	ID string `json:"id" gorm:"primarykey"`

	// fixed "Person"
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
	Devices string `json:"devices"`

	PublicKey *PublicKey `json:"publicKey" gorm:"foreignKey:Owner;references:ID"`

	Tag []string `json:"tag" gorm:"-"` // todo

	// what is the type?
	// not sure it's possibe the emojis
	Attachment []any `json:"attachment" gorm:"-"` // todo

	SharedInbox string     `json:"sharedInbox,omitempty"`
	Endpoint    *Endpoints `json:"endpoints" gorm:"-"`

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

func UserEndpoint(host string) *Endpoints {
	return &Endpoints{
		SharedInbox: "https://" + host + "/inbox",
	}
}

func NewUser(name, host string) *User {
	return &User{
		Context:           UserContext,
		ID:                utils.ParseActivitypubID(name, host),
		Type:              "Person",
		Following:         utils.ParseActivitypubID(name, host) + "/following",
		Followers:         utils.ParseActivitypubID(name, host) + "/followers",
		Inbox:             utils.ParseActivitypubID(name, host) + "/inbox",
		Outbox:            utils.ParseActivitypubID(name, host) + "/outbox",
		Featured:          utils.ParseActivitypubID(name, host) + "/collections/featured",
		FeaturedTags:      utils.ParseActivitypubID(name, host) + "/collections/tags",
		PreferredUsername: name,

		URL: utils.ParseProfileUrl(name, host),

		Published: utils.MicroSecondToRFC3339(utils.Now()),
		Devices:   utils.ParseActivitypubID(name, host) + "/collections/devices",
		PublicKey: NewPublicKey(utils.ParseActivitypubID(name, host)),
		Endpoint:  UserEndpoint(host),

		Icon: &Image{
			"Image",
			"image/png",
			"https://twimg.moonchan.xyz/media/GB8hT5vacAAK6wa?format=png&name=medium",
		},
	}
}

func (o *User) Autofill() {
	name, host := utils.ParseNameAndHost(o.ID)
	o.Context = UserContext
	o.Type = "Person"
	o.Following = o.ID + "/following"
	o.Followers = o.ID + "/followers"
	o.Inbox = o.ID + "/inbox"
	o.Outbox = o.ID + "/outbox"
	o.Featured = o.ID + "/collections/featured"
	o.FeaturedTags = o.ID + "/collections/tags"
	o.URL = utils.ParseProfileUrl(name, host)
	o.Devices = o.ID + "/collections/devices"
	o.Endpoint = UserEndpoint(host)
}

type IDType struct {
	ID   string `json:"@id"`
	Type string `json:"@type"`
}
