package convert

import (
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"

	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
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

func ToActivityPubFollow(o *core.LocalRelation) *activitypub.Follow {
	follow := &activitypub.Follow{
		ID:     o.ID,
		Actor:  o.Actor,
		Object: o.Object,
	}
	follow.Autofill()
	return follow
}
