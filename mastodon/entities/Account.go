package entities

import "github.com/Hana-ame/fedi-antenna/mastodon/entities/account"

type Account struct {
	// Type: String (cast from an integer, but not guaranteed to be a number)
	// Description: The account id.
	Id string `json:"id" gorm:"primarykey"`
	// Type: String
	// Description: The username of the account, not including domain.
	Username string `json:"username"`
	// Type: String
	// Description: The Webfinger account URI. Equal to username for local users, or username@domain for remote users.
	Acct string `json:"acct"`
	// Type: String (URL)
	// Description: The location of the user’s profile page.
	Url string `json:"url"`
	// Type: String
	// Description: URI of the status used for federation.
	Uri string `json:"uri" gorm:"index:,unique"` // activitypubID
	// Type: String
	// Description: The profile’s display name.
	DisplayName string `json:"display_name"`
	// Type: String (HTML)
	// Description: The profile’s bio or description.
	Note string `json:"note"`
	// Type: String (URL)
	// Description: An image icon that is shown next to statuses and in the profile.
	Avatar string `json:"avatar"`
	// Type: String (URL)
	// Description: A static version of the avatar. Equal to avatar if its value is a static image; different if avatar is an animated GIF.
	AvatarStatic string `json:"avatar_static"`
	// Type: String (URL)
	// Description: An image banner that is shown above the profile and in profile cards.
	Header string `json:"header"`
	// Type: String (URL)
	// Description: A static version of the header. Equal to header if its value is a static image; different if header is an animated GIF.
	HeaderStatic string `json:"header_static"`
	// Type: Boolean
	// Description: Whether the account manually approves follow requests.
	Locked bool `json:"locked"`
	// Type: Array of Field
	// Description: Additional metadata attached to a profile as name-value pairs.
	Fields []*account.Field `json:"fields" gorm:"serializer:json"`
	// Type: Array of CustomEmoji
	// Description: Custom emoji entities to be used when rendering the profile.
	Emojis []*CustomEmoji `json:"emojis" gorm:"-"`
	// Type: Boolean
	// Description: Indicates that the account may perform automated actions, may not be monitored, or identifies as a robot.
	Bot bool `json:"bot"`
	// Type: Boolean
	// Description: Indicates that the account represents a Group actor.
	Group bool `json:"group"`
	// Type: NULLABLE Boolean
	// Description: Whether the account has opted into discovery features such as the profile directory.
	Discoverable *bool `json:"discoverable"`
	// Type: NULLABLE Boolean
	// Description: Whether the local user has opted out of being indexed by search engines.
	Noindex *bool `json:"noindex,omitempty"`
	// Type: NULLABLE Account, or null if the profile is suspended.
	// Description: Indicates that the profile is currently inactive and that its user has moved to a new account.
	Moved *Account `json:"moved,omitempty" gorm:"-"`
	// Type: Boolean
	// Description: An extra attribute returned only when an account is suspended.
	Suspended bool `json:"suspended,omitempty"`
	// Type: Boolean
	// Description: An extra attribute returned only when an account is silenced. If true, indicates that the account should be hidden behind a warning screen.
	Limited bool `json:"limited,omitempty"`
	// Type: String (ISO 8601 Datetime)
	// Description: When the account was created.
	CreatedAt string `json:"created_at"`
	// Type: NULLABLE String (ISO 8601 Date), or null if no statuses
	// Description: When the most recent status was posted.
	LastStatusAt *string `json:"last_status_at"`
	// Type: Integer
	// Description: How many statuses are attached to this account.
	StatusesCount int `json:"statuses_count"`
	// Type: Integer
	// Description: The reported followers of this profile.
	FollowersCount int `json:"followers_count"`
	// Type: Integer
	// Description: The reported follows of this profile.
	FollowingCount int `json:"following_count"`
}
