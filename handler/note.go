package handler

import (
	"log"
	"strconv"

	"github.com/Hana-ame/fedi-antenna/activitypub/actions"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/core/utils/convert"
	mastodon "github.com/Hana-ame/fedi-antenna/mastodon/model"
)

func Note(actor string, o *mastodon.Status) error {
	// actor string,

	timestamp := utils.Now()
	// published := utils.TimestampToRFC3339(timestamp)
	name, host := utils.ParseNameAndHost(actor)
	id := utils.ParseStatusesID(name, host, strconv.Itoa(int(timestamp)))

	localNote := &core.LocalNote{
		ID:           id,
		AttributedTo: actor,
		Status:       o.Status,
		MediaIDs:     o.MediaIDs,
		InReplyToID:  o.InReplyToID,
		Sensitive:    o.Sensitive,
		SpoilerText:  o.SpoilerText,
		Visibility:   o.Visibility,
		Published:    timestamp,
	}

	note := convert.ToActivityPubNote(localNote)

	if err := dao.Create(localNote); err != nil {
		log.Println(err)
		return err
	}

	if err := actions.CreateNote(note); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
