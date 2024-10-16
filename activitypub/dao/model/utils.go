package model

// todo
const (
	TypeFollow = "Follow"
	TypeBlock  = "Block"
	TypeUndo   = "Undo"
	TypeAccept = "Accept"
	TypeReject = "Reject"

	TypeCreate = "Create"
	TypeDelete = "Delete"

	TypePerson   = "Person"
	TypeNote     = "Note"
	TypeLike     = "Like"
	TypeAnnounce = "Announce"
	TypeMention  = "Mention"
	TypeImage    = "Image"

	TypeUnknown = "Unknown"
)

type IDType struct {
	ID   string `json:"@id"`
	Type string `json:"@type"`
}
