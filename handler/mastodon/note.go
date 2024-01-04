package handler

import (
	"strconv"

	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/core/actions"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/model"
)

func Note(actor string, o *model.Status) error {
	// actor string,
	
	// content string,
	var visibility int
	switch o.Visibility {
	case "public":
		visibility = activitypub.VisiblityPublic
	case "unlisted":
		visibility = activitypub.VisiblityUnlisted
	case "private":
		visibility = activitypub.VisiblityPrivate
	case "direct":
		visibility = activitypub.VisiblityDirect
	}

	timestamp := utils.Now()
	published := utils.TimestampToRFC3339(timestamp)
	name, host := utils.ParseNameAndHost(actor)

	note := &activitypub.Note{
		ID:          utils.ParseStatusesID(name, host, strconv.Itoa(int(timestamp))),
		Summary:     utils.ParseStringToPointer(o.SpoilerText, true),
		Content:     o.Status,
		Visibility:  visibility,
		InReplyTo:   utils.ParseStringToPointer(o.InReplyToID, true),
		AttributeTo: actor,
		Published:   published,
	}
	note.Autofill()

	if err := dao.Create(note); err != nil {
		return err
	}

	err := actions.CreateNote(note)

	return err
}
