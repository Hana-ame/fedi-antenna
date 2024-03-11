package handler

import (
	"strconv"

	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/db"
	"github.com/Hana-ame/fedi-antenna/mastodon/controller/statuses/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/dao"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

// todo: poll
// todo: image list
func Post_a_new_status(
	actor, IdempotencyKey string, o *model.Post_a_new_status,
) (*entities.Status, error) {

	timestamp := utils.Timestamp(false)
	name, host := utils.ActivitypubID2NameAndHost(actor)

	acct := &entities.Account{
		Uri: actor,
	}

	tx := db.Begin()

	if err := dao.Read(tx, acct); err != nil {
		logE(err)
		return nil, err
	}

	inReplyToAccountId := new(string)
	if o.InReplyToId != nil {
		replyto := &entities.Status{Id: *o.InReplyToId}
		if err := dao.ReadStatuses(tx, replyto); err != nil {
			logE(err)
			return nil, err
		}
		inReplyToAccountId = &replyto.Account.Id
	}

	mediaAttachments := make([]*entities.MediaAttachment, len(o.MediaIds))
	for i := range o.MediaIds {
		mediaAttachment := &entities.MediaAttachment{PreviewUrl: o.MediaIds[i]}
		db.Read(tx, mediaAttachment) // do not process error
		mediaAttachments[i] = mediaAttachment
	}

	status := &entities.Status{
		// Type: String (cast from an integer but not guaranteed to be a number)
		// Description: ID of the status in the database.
		Id: strconv.Itoa(int(timestamp)),
		// Type: String
		// Description: URI of the status used for federation.
		Uri: utils.NameHostTimestampToStatusesUri(name, host, strconv.Itoa(int(timestamp))),
		// Type: String (ISO 8601 Datetime)
		// Description: The date when this status was created.
		CreatedAt: utils.TimestampToRFC3339(timestamp),
		// Type: Account
		// Description: The account that authored this status.
		AttributedTo: actor,
		Account:      acct,
		// Type: String (HTML)
		// Description: HTML-encoded status content.
		Content: o.Status,
		// Type: String (Enumerable oneOf)
		// Description: Visibility of this status.
		Visibility: o.Visibility,
		// Type: Boolean
		// Description: Is this status marked as sensitive content?
		Sensitive: o.Sensitive,
		// Type: String
		// Description: Subject or summary line, below which status content is collapsed until expanded.
		SpoilerText: o.SpoilerText,
		// Type: Array of MediaAttachment
		// Description: Media that is attached to this status.
		MediaAttachments: mediaAttachments,
		// Type: Hash
		// Description: The application used to post this status.
		// Application *status.Application `json:"application,omitempty" gorm:"serializer:json"`

		// Mentions []*status.Mention `json:"mentions" gorm:"-"`
		// Tags []*status.Tag `json:"tags" gorm:"-"`
		// Emojis []*CustomEmoji `json:"emojis" gorm:"-"`
		// ReblogsCount int `json:"reblogs_count"`
		// FavouritesCount int `json:"favourites_count"`
		// RepliesCount int `json:"replies_count"`

		// Type: NULLABLE String (URL) or null
		// Description: A link to the status’s HTML representation.
		Url: utils.ParseStringToPointer(utils.ParseStatusesURL(name, host, strconv.Itoa(int(timestamp))), true),
		// Type: NULLABLE String (cast from an integer but not guaranteed to be a number) or null
		// Description: ID of the status being replied to.
		InReplyToId: o.InReplyToId,
		// Type: NULLABLE String (cast from an integer but not guaranteed to be a number) or null
		// Description: ID of the account that authored the status being replied to.
		InReplyToAccountId: inReplyToAccountId,

		// Reblog *Status `json:"reblog"`
		// Poll *Poll `json:"poll" gorm:"foreignKey:Id;references:Id"`
		// Card *PreviewCard `json:"card" gorm:"serializer:json"`
		// Language *string `json:"language"`
		// Text *string `json:"text"`
		// EditedAt *string `json:"edited_at"`
	}

	if err := dao.CreateStatus(tx, status); err != nil {
		logE(err)
		tx.Rollback()
		return status, err
	}

	tx.Commit()

	return status, tx.Error
}

func Delete_a_status(id string, actor string) (*entities.Status, error) {

	status := &entities.Status{
		Id: id,
	}

	tx := db.Begin()

	if err := dao.ReadStatuses(tx, status); err != nil {
		logE(err)
		return nil, err
	}

	if err := dao.DeleteStatus(tx, status); err != nil {
		logE(err)
		return nil, err
	}

	tx.Commit()

	return status, tx.Error
}
