package handler

import (
	"log"
	"strconv"

	"github.com/Hana-ame/fedi-antenna/actions"
	c "github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	mastodon "github.com/Hana-ame/fedi-antenna/mastodon/controller/statuses/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

// todo: poll
// todo: image list
func Post_a_new_status(actor, IdempotencyKey string, o *mastodon.Post_a_new_status) (*entities.Status, error) {
	// actor string,

	timestamp := utils.Now()
	name, host := utils.ParseNameAndHost(actor)

	account := &entities.Account{
		Uri: actor,
	}
	if err := dao.Read(account); err != nil {
		log.Println(err)
		return nil, err
	}

	var inReplyToAccountId *string
	if o.InReplyToId != nil {
		replyto := &entities.Status{Id: *o.InReplyToId}
		if err := dao.Read(replyto); err != nil {
			log.Printf("%s", err.Error())
			return nil, err
		}
		inReplyToAccountId = &replyto.AttributedTo
	}

	status := &entities.Status{
		// Type: String (cast from an integer but not guaranteed to be a number)
		// Description: ID of the status in the database.
		Id: strconv.Itoa(int(timestamp)),
		// Type: String
		// Description: URI of the status used for federation.
		Uri: utils.ParseStatusesUri(name, host, strconv.Itoa(int(timestamp))),
		// Type: String (ISO 8601 Datetime)
		// Description: The date when this status was created.
		CreatedAt: utils.TimestampToRFC3339(timestamp),
		// Type: Account
		// Description: The account that authored this status.
		AttributedTo: actor,
		Account:      account,
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
		// MediaAttachments []*MediaAttachment `json:"media_attachments" gorm:"many2many:status_mediaattachments;"`
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

	err := c.CreateStatus(status)

	if err == nil {
		actions.CreateNote(convert.ToActivityPubNote(status))
	}

	return status, err
}

func Delete_a_status(id string, actor string) (*entities.Status, error) {

	status := &entities.Status{
		Id:           id,
		AttributedTo: actor,
	}

	err := c.DeleteStatus(status)

	if err == nil {

	}

	return status, err
}
