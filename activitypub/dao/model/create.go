package model

import "github.com/Hana-ame/fedi-antenna/utils"

type Create struct {
	Context   any    `json:"@context" gorm:"-"`
	ID        string `json:"id" gorm:"primarykey"`
	Type      string `json:"type" gorm:"-"`
	Actor     string `json:"actor"`
	Published string `json:"published"`

	// public, social, private, direct
	// Visiblity int      `json:"-"`
	To []string `json:"to" gorm:"-"`
	Cc []string `json:"cc" gorm:"-"`

	ObjectID string    `json:"-"`
	Object   Creatable `json:"object" gorm:"-"`

	Signature *Signature `json:"signature,omitempty" gorm:"-"`
}

const (
	VisiblityPublic          = "public"
	VisiblityUnlisted        = "unlisted"
	VisiblityPrivate         = "private"
	VisiblityDirect          = "direct"
	VisiblityIsolatePublic   = "isolate,public"
	VisiblityIsolateUnlisted = "isolate,unlisted"
	VisiblityIsolatePrivate  = "isolate,private"
	VisiblityIsolateDirect   = "isolate,direct"
)

var (
	EndpointPublic = []string{"https://www.w3.org/ns/activitystreams#Public"}
)

type Creatable interface {
	ClearContext()
	GetID() string
	GetTo() []string
	GetCc() []string
	GetActor() string
	// GetPublished() string
}

var CreateContext = []any{
	"https://www.w3.org/ns/activitystreams",
	utils.NewMapFromKV([]*utils.KV{
		{Key: "ostatus", Value: "http://ostatus.org#"},
		{Key: "atomUri", Value: "ostatus:atomUri"},
		{Key: "inReplyToAtomUri", Value: "ostatus:inReplyToAtomUri"},
		{Key: "conversation", Value: "ostatus:conversation"},
		{Key: "sensitive", Value: "as:sensitive"},
		{Key: "toot", Value: "http://joinmastodon.org/ns#"},
		{Key: "votersCount", Value: "toot:votersCount"},
	}),
}

// only for note.
func (o *Create) Autofill() {
	o.Context = CreateContext
	o.Object.ClearContext()
	o.Type = TypeCreate
	o.Actor = o.Object.GetActor()
	o.ObjectID = o.Object.GetID()
	o.ID = o.Object.GetID() + "/activity" // only for note.
	o.To = o.Object.GetTo()
	o.Cc = o.Object.GetCc()

	// todo
	//Signature
}
