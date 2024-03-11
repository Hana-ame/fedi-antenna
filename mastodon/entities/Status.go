package entities

import "github.com/Hana-ame/fedi-antenna/mastodon/entities/status"

type Status struct {
	// Type: String (cast from an integer but not guaranteed to be a number)
	// Description: ID of the status in the database.
	Id string `json:"id" gorm:"primaryKey"`
	// Type: String
	// Description: URI of the status used for federation.
	Uri string `json:"uri"` // activitypubID
	// Type: String (ISO 8601 Datetime)
	// Description: The date when this status was created.
	CreatedAt string `json:"created_at"`
	// Type: Account
	// Description: The account that authored this status.
	AttributedTo string   `json:"attribute_to"`
	Account      *Account `json:"account" gorm:"foreignKey:AttributedTo;references:Uri"`
	// Type: String (HTML)
	// Description: HTML-encoded status content.
	Content string `json:"content"`
	// Type: String (Enumerable oneOf)
	// Description: Visibility of this status.
	Visibility string `json:"visibility"`
	// Type: Boolean
	// Description: Is this status marked as sensitive content?
	Sensitive bool `json:"sensitive"`
	// Type: String
	// Description: Subject or summary line, below which status content is collapsed until expanded.
	SpoilerText string `json:"spoiler_text"`
	// Type: Array of MediaAttachment
	// Description: Media that is attached to this status.
	MediaAttachments []*MediaAttachment `json:"media_attachments" gorm:"many2many:status_mediaattachments;"`
	// Type: Hash
	// Description: The application used to post this status.
	Application *status.Application `json:"application,omitempty" gorm:"serializer:json"`
	// Type: Array of Status::Mention
	// Description: Mentions of users within the status content.
	Mentions []*status.Mention `json:"mentions" gorm:"-"`
	// Type: Array of Status::Tag
	// Description: Hashtags used within the status content.
	Tags []*status.Tag `json:"tags" gorm:"-"`
	// Type: Array of CustomEmoji
	// Description: Custom emoji to be used when rendering status content.
	Emojis []*CustomEmoji `json:"emojis" gorm:"-"`
	// Type: Integer
	// Description: How many boosts this status has received.
	ReblogsCount int `json:"reblogs_count"`
	// Type: Integer
	// Description: How many favourites this status has received.
	FavouritesCount int `json:"favourites_count"`
	// Type: Integer
	// Description: How many replies this status has received.
	RepliesCount int `json:"replies_count"`
	// Type: NULLABLE String (URL) or null
	// Description: A link to the status’s HTML representation.
	Url *string `json:"url"`
	// Type: NULLABLE String (cast from an integer but not guaranteed to be a number) or null
	// Description: ID of the status being replied to.
	InReplyToId *string `json:"in_reply_to_id"`
	// Type: NULLABLE String (cast from an integer but not guaranteed to be a number) or null
	// Description: ID of the account that authored the status being replied to.
	InReplyToAccountId *string `json:"in_reply_to_account_id"`
	// Type: NULLABLE Status or null
	// Description: The status being reblogged.
	Reblog *Status `json:"reblog" gorm:"-"`
	// Type: NULLABLE Poll or null
	// Description: The poll attached to the status.
	Poll *Poll `json:"poll" gorm:"foreignKey:Id;references:Id"`
	// Type: NULLABLE PreviewCard or null
	// Description: Preview card for links included within status content.
	Card *PreviewCard `json:"card" gorm:"serializer:json"`
	// Type: NULLABLE String (ISO 639 Part 1 two-letter language code) or null
	// Description: Primary language of this status.
	Language *string `json:"language"`
	// Type: NULLABLE String or null
	// Description: Plain-text source of a status. Returned instead of content when status is deleted, so the user may redraft from the source text without the client having to reverse-engineer the original text from the HTML content.
	Text *string `json:"text"`
	// Type: NULLABLE String (ISO 8601 Datetime)
	// Description: Timestamp of when the status was last edited.
	EditedAt *string `json:"edited_at"`
	// Type: Boolean
	// Description: If the current token has an authorized user: Have you favourited this status?
	Favourited bool `json:"favourited,omitempty" gorm:"-"`
	// Type: Boolean
	// Description: If the current token has an authorized user: Have you boosted this status?
	Reblogged bool `json:"reblogged,omitempty" gorm:"-"`
	// Type: Boolean
	// Description: If the current token has an authorized user: Have you muted notifications for this status’s conversation?
	Muted bool `json:"muted,omitempty" gorm:"-"`
	// Type: Boolean
	// Description: If the current token has an authorized user: Have you bookmarked this status?
	Bookmarked bool `json:"bookmarked,omitempty" gorm:"-"`
	// Type: Boolean
	// Description: If the current token has an authorized user: Have you pinned this status? Only appears if the status is pinnable.
	Pinned bool `json:"pinned,omitempty" gorm:"-"`
	// Type: Array of FilterResult
	// Description: If the current token has an authorized user: The filter and keywords that matched this status.
	Filtered []FilterResult `json:"filtered,omitempty" gorm:"-"`

	DeletedAt int64 `json:"-"`
}
