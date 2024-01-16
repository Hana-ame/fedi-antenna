package model

import "github.com/Hana-ame/fedi-antenna/core/utils"

type Announce struct {
	Context any    `json:"@context,omitempty" gorm:"-"`
	ID      string `json:"id" gorm:"primarykey"`
	Type    string `json:"type" gorm:"-"`

	Actor string `json:"actor"`

	Published string `json:"published"`

	Visibility string `json:"-"`
	To         []string
	Cc         []string

	// id of statuses
	Object string `json:"object"`
}

var AnnounceContext = "https://www.w3.org/ns/activitystreams"

func (o *Announce) Autofill() {
	o.Context = AnnounceContext
	o.Type = TypeAnnounce

	name, host, _ := utils.ParseStatusesIDToNameHostTimestamp(o.ID)
	switch o.Visibility {
	case VisiblityPublic:
		o.To = EndpointPublic
		o.Cc = []string{endpointFollower(name, host)}
	case VisiblityUnlisted:
		o.To = []string{endpointFollower(name, host)}
		o.Cc = EndpointPublic
	case VisiblityPrivate:
		o.To = []string{endpointFollower(name, host)}
	case VisiblityDirect:
	default:
		o.Visibility = utils.ParseVisibility(o.To, o.Cc)
	}
}

func (o *Announce) ClearContext() {
	o.Context = nil
}
func (o *Announce) GetID() string {
	return o.ID
}
func (o *Announce) GetActor() string {
	return o.Actor
}
func (o *Announce) GetEndpoint() string {
	return "https://www.w3.org/ns/activitystreams#Public"
}
