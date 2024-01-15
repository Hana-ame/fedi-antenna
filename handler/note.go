package handler

import (
	"log"
	"strconv"

	"github.com/Hana-ame/fedi-antenna/activitypub/actions"
	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	mastodon "github.com/Hana-ame/fedi-antenna/mastodon/controller/statuses/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Post_a_new_status(actor, IdempotencyKey string, o *mastodon.Post_a_new_status) (*entities.Status, error) {
	// actor string,

	timestamp := utils.Now()
	// published := utils.TimestampToRFC3339(timestamp)
	name, host := utils.ParseNameAndHost(actor)
	id := utils.ParseStatusesID(name, host, strconv.Itoa(int(timestamp)))

	localNote := &core.LocalNote{
		ID:           id,
		AttributedTo: actor,
		Status:       o.Status,
		MediaIDs:     o.MediaIds,
		InReplyToID:  o.InReplyToId,
		Sensitive:    o.Sensitive,
		SpoilerText:  o.SpoilerText,
		Visibility:   o.Visibility,
		Published:    timestamp,
	}

	if err := dao.Create(localNote); err != nil {
		log.Println(err)
		return nil, err
	}

	// activitypub
	note := convert.ToActivityPubNote(localNote)

	if err := actions.CreateNote(note, nil, false); err != nil {
		log.Println(err)
		return nil, err
	}

	// mastodon
	status := convert.LocalNoteToMastodonStatus(localNote)

	return status, nil
}

func Delete_a_status(id string, actor string) (*entities.Status, error) {
	return nil, nil
}
