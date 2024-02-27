package model

type Like struct {
	Context any    `json:"@context,omitempty" gorm:"-"`
	ID      string `json:"id" gorm:"primarykey"`
	Type    string `json:"type" gorm:"-"`

	Actor string `json:"actor"`

	// id of statuses
	Object string `json:"object"`
}

var LikeContext = "https://www.w3.org/ns/activitystreams"

func (o *Like) Autofill() {
	o.Context = LikeContext
	o.Type = TypeLike
}

func (o *Like) ClearContext() {
	o.Context = nil
}
func (o *Like) GetID() string {
	return o.ID
}
func (o *Like) GetActor() string {
	return o.Actor
}
func (o *Like) GetEndpoint() string {
	return "https://www.w3.org/ns/activitystreams#Public"
}

