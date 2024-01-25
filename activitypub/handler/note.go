package handler

import (
	"strconv"

	tools "github.com/Hana-ame/fedi-antenna/Tools"
	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities/status"
)

// type: note
func CreateNote(o *orderedmap.OrderedMap) error {
	timestamp := utils.Now()
	oapplication, ok := o.GetOrDefault("application", orderedmap.New()).(*orderedmap.OrderedMap)
	var application *status.Application
	if ok && oapplication != nil {
		application = &status.Application{
			Name:    oapplication.GetOrDefault("name", "").(string),
			Website: utils.ParseStringToPointer(oapplication.GetOrDefault("website", "").(string), true),
		}
	}
	n := &entities.Status{
		// Type: String (cast from an integer but not guaranteed to be a number)
		// Description: ID of the status in the database.
		Id: strconv.Itoa(int(timestamp)),
		// Type: String
		// Description: URI of the status used for federation.
		Uri: o.GetOrDefault("id", "").(string),
		// Type: String (ISO 8601 Datetime)
		// Description: The date when this status was created.
		CreatedAt: o.GetOrDefault("published", "").(string),
		// Type: Account
		// Description: The account that authored this status.
		AttributedTo: o.GetOrDefault("attributedTo", "").(string),
		// Account      *Account `json:"account" gorm:"foreignKey:AttributedTo;references:Uri"`
		// Type: String (HTML)
		// Description: HTML-encoded status content.
		Content: o.GetOrDefault("content", "").(string),
		// Type: String (Enumerable oneOf)
		// Description: Visibility of this status.
		Visibility: utils.ParseVisibility(
			tools.ParseSliceToStringSlice(o.Get("to")),
			tools.ParseSliceToStringSlice(o.Get("cc")),
		),
		// Type: Boolean
		// Description: Is this status marked as sensitive content?
		Sensitive: o.GetOrDefault("sensitive", false).(bool),
		// Type: String
		// Description: Subject or summary line, below which status content is collapsed until expanded.
		SpoilerText: o.GetOrDefault("summary", "").(string),
		// Type: Array of MediaAttachment
		// Description: Media that is attached to this status.
		MediaAttachments: []*entities.MediaAttachment{},
		// Type: Hash
		// Description: The application used to post this status.
		Application: application,
		// Type: Array of Status::Mention
		// Description: Mentions of users within the status content.
		Mentions: []*status.Mention{},
		// Tags []*status.Tag `json:"tags" gorm:"-"`
		// Emojis []*CustomEmoji `json:"emojis" gorm:"-"`
		// ReblogsCount int `json:"reblogs_count"`
		// FavouritesCount int `json:"favourites_count"`
		// RepliesCount int `json:"replies_count"`
		Url: utils.ParseStringToPointer(o.GetOrDefault("url", "").(string), true),
		// Type: NULLABLE String (cast from an integer but not guaranteed to be a number) or null
		// Description: ID of the status being replied to.
		InReplyToId: utils.ParseStringToPointer(o.GetOrDefault("inReplyTo", "").(string), true),
		// Type: NULLABLE String (cast from an integer but not guaranteed to be a number) or null
		// Description: ID of the account that authored the status being replied to.
		InReplyToAccountId: nil,
		Reblog:             nil,
		// Type: NULLABLE Poll or null
		// Description: The poll attached to the status.
		Poll: nil,
		// Type: NULLABLE PreviewCard or null
		// Description: Preview card for links included within status content.
		Card: nil,
		// Type: NULLABLE String (ISO 639 Part 1 two-letter language code) or null
		// Description: Primary language of this status.
		Language: nil,
		// Type: NULLABLE String or null
		// Description: Plain-text source of a status. Returned instead of content when status is deleted, so the user may redraft from the source text without the client having to reverse-engineer the original text from the HTML content.
		Text: nil,
		// Type: NULLABLE String (ISO 8601 Datetime)
		// Description: Timestamp of when the status was last edited.
		EditedAt: nil,
		// Favourited bool `json:"favourited,omitempty" gorm:"-"`
		// Reblogged bool `json:"reblogged,omitempty" gorm:"-"`
		// Muted bool `json:"muted,omitempty" gorm:"-"`
		// Bookmarked bool `json:"bookmarked,omitempty" gorm:"-"`
		// Pinned bool `json:"pinned,omitempty" gorm:"-"`
		// Filtered []FilterResult `json:"filtered,omitempty" gorm:"-"`
	}
	err := core.CreateStatus(n)
	return err
}

func DeleteNote(id string) error {
	status := &entities.Status{
		Uri: id,
	}
	err := core.DeleteStatus(status)
	return err
}
