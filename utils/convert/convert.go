package convert

import (
	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"

	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
)

func ToActivityPubNote(o *core.LocalNote) *activitypub.Note {
	note := &activitypub.Note{
		ID:          o.ID,
		Summary:     utils.ParseStringToPointer(o.SpoilerText, true),
		Content:     o.Status,
		Visibility:  o.Visibility,
		InReplyTo:   utils.ParseStringToPointer(o.InReplyToID, true),
		AttributeTo: o.AttributedTo,
		Published:   utils.TimestampToRFC3339(o.Published),
	}
	note.Autofill()
	return note
}

func ToActivityPubUser(o *core.LocalUser) *activitypub.User {
	// name, host := utils.ParseNameAndHost(o.ID)
	icon := &activitypub.Image{
		URL: o.IconURL,
	}
	dao.Read(icon)

	user := &activitypub.User{
		ID:                        o.ID,
		Name:                      o.Name,
		Summary:                   o.Summary,
		ManuallyApprovesFollowers: o.ManuallyApprovesFollowers,
		Discoverable:              o.Discoverable,
		Published:                 utils.TimestampToRFC3339(o.Published),
		AlsoKnownAs:               o.AlsoKnownAs,
		PublicKey:                 &activitypub.PublicKey{ID: o.ID + "#main-key", Owner: o.ID, PublicKeyPem: o.PublicKeyPem},
		Tag:                       nil,
		Attachment:                nil,
		SharedInbox:               "",
		Endpoint:                  nil,
		IconURL:                   o.IconURL,
		Icon:                      icon,
	}
	user.Autofill()
	return user
}
