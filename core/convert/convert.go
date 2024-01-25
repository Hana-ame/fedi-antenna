package convert

import (
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"

	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
)

func ToActivityPubNote(o *entities.Status) *activitypub.Note {
	note := &activitypub.Note{
		ID:           o.Uri,
		Summary:      utils.ParseStringToPointer(o.SpoilerText, true),
		Content:      o.Content,
		Visibility:   o.Visibility,
		InReplyTo:    o.InReplyToId,
		AttributedTo: o.AttributedTo,
		Published:    o.CreatedAt,
	}
	note.Autofill()
	return note
}

func ToActivityPubFollow(o *core.LocalRelation) *activitypub.Follow {
	follow := &activitypub.Follow{
		ID:     o.ID,
		Actor:  o.Actor,
		Object: o.Object,
	}
	follow.Autofill()
	return follow
}
