package convert

import (
	"log"
	"strconv"

	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities/status"
)

// this function may substanced by local2activitypub and activitypub2mastodon?
// func ToMastodonAccount(lu *core.LocalUser) *entities.Account {
// 	name, host := utils.ParseNameAndHost(lu.ID)
// 	account := &entities.Account{
// 		// Type: String (cast from an integer, but not guaranteed to be a number)
// 		// Description: The account id.
// 		Id: utils.TimestampToRFC3339(lu.Published),
// 		// Type: String
// 		// Description: The username of the account, not including domain.
// 		Username: lu.PreferredUsername,
// 		// Type: String
// 		// Description: The Webfinger account URI. Equal to username for local users, or username@domain for remote users.
// 		Acct: utils.ParseAcctStr(name, host),
// 		// Type: String (URL)
// 		// Description: The location of the user’s profile page.
// 		Url: utils.ParseProfileUrl(name, host),
// 		// Type: String
// 		// Description: The profile’s display name.
// 		DisplayName: lu.Name,
// 		// Type: String (HTML)
// 		// Description: The profile’s bio or description.
// 		Note: lu.Summary,
// 		// Type: String (URL)
// 		// Description: An image icon that is shown next to statuses and in the profile.
// 		Avatar: lu.IconURL,
// 		// Type: String (URL)
// 		// Description: A static version of the avatar. Equal to avatar if its value is a static image; different if avatar is an animated GIF.
// 		AvatarStatic: lu.IconURL,
// 		// Type: String (URL)
// 		// Description: An image banner that is shown above the profile and in profile cards.
// 		Header: lu.ImageURL,
// 		// Type: String (URL)
// 		// Description: A static version of the header. Equal to header if its value is a static image; different if header is an animated GIF.
// 		HeaderStatic: lu.ImageURL,
// 		// Type: Boolean
// 		// Description: Whether the account manually approves follow requests.
// 		Locked: lu.ManuallyApprovesFollowers,
// 		// Type: Array of Field
// 		// Description: Additional metadata attached to a profile as name-value pairs.
// 		Fields: []*account.Field{},
// 		// Type: Array of CustomEmoji
// 		// Description: Custom emoji entities to be used when rendering the profile.
// 		Emojis: []*entities.CustomEmoji{},
// 		// Type: Boolean
// 		// Description: Indicates that the account may perform automated actions, may not be monitored, or identifies as a robot.
// 		Bot: lu.IsBot,
// 		// Type: Boolean
// 		// Description: Indicates that the account represents a Group actor.
// 		Group: lu.IsGroup,
// 		// Type: NULLABLE Boolean
// 		// Description: Whether the account has opted into discovery features such as the profile directory.
// 		Discoverable: utils.ParseBoolToPointer(lu.IsDiscoverable, true),
// 		// Type: NULLABLE Boolean
// 		// Description: Whether the local user has opted out of being indexed by search engines.
// 		Noindex: utils.ParseBoolToPointer(lu.IsNoindex, true),
// 		// Type: NULLABLE Account, or null if the profile is suspended.
// 		// Description: Indicates that the profile is currently inactive and that its user has moved to a new account.
// 		Moved: nil,
// 		// Type: Boolean
// 		// Description: An extra attribute returned only when an account is suspended.
// 		Suspended: lu.IsSuspended,
// 		// Type: Boolean
// 		// Description: An extra attribute returned only when an account is silenced. If true, indicates that the account should be hidden behind a warning screen.
// 		Limited: lu.IsLimited,
// 		// Type: String (ISO 8601 Datetime)
// 		// Description: When the account was created.
// 		CreatedAt: utils.TimestampToRFC3339(lu.Published),
// 		// Type: NULLABLE String (ISO 8601 Date), or null if no statuses
// 		// Description: When the most recent status was posted.
// 		LastStatusAt: utils.ParseStringToPointer(utils.TimestampToRFC3339(lu.LastSeenAt), true),
// 		// Type: Integer
// 		// Description: How many statuses are attached to this account.
// 		StatusesCount: lu.StatusesCount,
// 		// Type: Integer
// 		// Description: The reported followers of this profile.
// 		FollowersCount: lu.FollowersCount,
// 		// Type: Integer
// 		// Description: The reported follows of this profile.
// 		FollowingCount: lu.FollowingCount,
// 	}
// 	return account
// }

func ToMastodonStatus_Deprecated(note *core.LocalNote, notify *core.LocalNotify) *entities.Status {
	if note != nil {
		// return ToMastodonStatus(note)
	} else if notify != nil {
		return ToMastodonReblog(notify)
	} else {
		return nil
	}
	return nil
}

func ToMastodonReblog(ln *core.LocalNotify) *entities.Status {
	name, host, timestampString := utils.ParseStatusesUriToNameHostTimestamp(ln.ID)
	timestamp, err := strconv.Atoi(timestampString)
	if err != nil {
		log.Printf("%s", err.Error())
		return nil
	}
	lu := &core.LocalUser{
		AccountID: utils.ParseActivitypubID(name, host),
	}
	if err := dao.Read(lu); err != nil {
		log.Printf("%s", err.Error())
		return nil
	}
	acct := &entities.Account{
		Uri: utils.ParseActivitypubID(name, host),
	}
	if err := dao.Read(acct); err != nil {
		log.Printf("%s", err.Error())
		return nil
	}

	reblog := &entities.Status{
		Uri: ln.Object,
	}
	if err := dao.Read(reblog); err != nil {
		log.Printf("%s", err.Error())
		return nil
	}

	status := &entities.Status{
		// Type: String (cast from an integer but not guaranteed to be a number)
		// Description: ID of the status in the database.
		Id: timestampString,
		// Type: String
		// Description: URI of the status used for federation.
		Uri: ln.ID,
		// Type: String (ISO 8601 Datetime)
		// Description: The date when this status was created.
		CreatedAt: utils.TimestampToRFC3339(int64(timestamp)),
		// Type: Account
		// Description: The account that authored this status.
		Account: acct,
		// Type: String (HTML)
		// Description: HTML-encoded status content.
		Content: "",
		// Type: String (Enumerable oneOf)
		// Description: Visibility of this status.
		Visibility: ln.Visibility,
		// Type: Boolean
		// Description: Is this status marked as sensitive content?
		Sensitive: false,
		// Type: String
		// Description: Subject or summary line, below which status content is collapsed until expanded.
		SpoilerText: "",
		// Type: Array of MediaAttachment
		// Description: Media that is attached to this status.
		MediaAttachments: []*entities.MediaAttachment{},
		// Type: Hash
		// Description: The application used to post this status.
		Application: nil,
		// Type: Array of Status::Mention
		// Description: Mentions of users within the status content.
		Mentions: []*status.Mention{},
		// Type: Array of Status::Tag
		// Description: Hashtags used within the status content.
		Tags: []*status.Tag{},
		// Type: Array of CustomEmoji
		// Description: Custom emoji to be used when rendering status content.
		Emojis: []*entities.CustomEmoji{},
		// Type: Integer
		// Description: How many boosts this status has received.
		ReblogsCount: 0,
		// Type: Integer
		// Description: How many favourites this status has received.
		FavouritesCount: 0,
		// Type: Integer
		// Description: How many replies this status has received.
		RepliesCount: 0,
		// Type: NULLABLE String (URL) or null
		// Description: A link to the status’s HTML representation.
		Url: utils.ParseStringToPointer(ln.ID, true),
		// Type: NULLABLE String (cast from an integer but not guaranteed to be a number) or null
		// Description: ID of the status being replied to.
		InReplyToId: nil,
		// Type: NULLABLE String (cast from an integer but not guaranteed to be a number) or null
		// Description: ID of the account that authored the status being replied to.
		InReplyToAccountId: nil,
		// Type: NULLABLE Status or null
		// Description: The status being reblogged.
		Reblog: reblog,
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
		// Type: Boolean
		// Description: If the current token has an authorized user: Have you favourited this status?
		Favourited: false,
		// Type: Boolean
		// Description: If the current token has an authorized user: Have you boosted this status?
		Reblogged: false,
		// Type: Boolean
		// Description: If the current token has an authorized user: Have you muted notifications for this status’s conversation?
		Muted: false,
		// Type: Boolean
		// Description: If the current token has an authorized user: Have you bookmarked this status?
		Bookmarked: false,
		// Type: Boolean
		// Description: If the current token has an authorized user: Have you pinned this status? Only appears if the status is pinnable.
		Pinned: false,
		// Type: Array of FilterResult
		// Description: If the current token has an authorized user: The filter and keywords that matched this status.
		Filtered: nil,
	}
	return status
}

// not used cuz use status to save at local insted
// func ToMastodonStatus(ln *core.LocalNote) *entities.Status {
// 	name, host := utils.ParseNameAndHost(ln.AttributedTo)
// 	lu := &core.LocalUser{
// 		AccountID: utils.ParseActivitypubID(name, host),
// 	}
// 	if err := dao.Read(lu); err != nil {
// 		log.Printf("%s", err.Error())
// 		return nil
// 	}
// 	acct := &entities.Account{
// 		Uri: utils.ParseActivitypubID(name, host),
// 	}
// 	if err := dao.Read(acct); err != nil {
// 		log.Printf("%s", err.Error())
// 		return nil
// 	}

// 	status := &entities.Status{
// 		// Type: String (cast from an integer but not guaranteed to be a number)
// 		// Description: ID of the status in the database.
// 		Id: strconv.Itoa(int(ln.Published)),
// 		// Type: String
// 		// Description: URI of the status used for federation.
// 		Uri: ln.ID,
// 		// Type: String (ISO 8601 Datetime)
// 		// Description: The date when this status was created.
// 		CreatedAt: utils.TimestampToRFC3339(ln.Published),
// 		// Type: Account
// 		// Description: The account that authored this status.
// 		Account: acct,
// 		// Type: String (HTML)
// 		// Description: HTML-encoded status content.
// 		Content: ln.Status,
// 		// Type: String (Enumerable oneOf)
// 		// Description: Visibility of this status.
// 		Visibility: ln.Visibility,
// 		// Type: Boolean
// 		// Description: Is this status marked as sensitive content?
// 		Sensitive: ln.Sensitive,
// 		// Type: String
// 		// Description: Subject or summary line, below which status content is collapsed until expanded.
// 		SpoilerText: ln.SpoilerText,
// 		// Type: Array of MediaAttachment
// 		// Description: Media that is attached to this status.
// 		MediaAttachments: []*entities.MediaAttachment{},
// 		// Type: Hash
// 		// Description: The application used to post this status.
// 		Application: nil,
// 		// Type: Array of Status::Mention
// 		// Description: Mentions of users within the status content.
// 		Mentions: []*status.Mention{},
// 		// Type: Array of Status::Tag
// 		// Description: Hashtags used within the status content.
// 		Tags: []*status.Tag{},
// 		// Type: Array of CustomEmoji
// 		// Description: Custom emoji to be used when rendering status content.
// 		Emojis: []*entities.CustomEmoji{},
// 		// Type: Integer
// 		// Description: How many boosts this status has received.
// 		ReblogsCount: ln.ReblogsCount,
// 		// Type: Integer
// 		// Description: How many favourites this status has received.
// 		FavouritesCount: ln.FavouritesCount,
// 		// Type: Integer
// 		// Description: How many replies this status has received.
// 		RepliesCount: ln.RepliesCount,
// 		// Type: NULLABLE String (URL) or null
// 		// Description: A link to the status’s HTML representation.
// 		Url: utils.ParseStringToPointer(utils.ParseStatusesURL(name, host, ln.ID), true),
// 		// Type: NULLABLE String (cast from an integer but not guaranteed to be a number) or null
// 		// Description: ID of the status being replied to.
// 		InReplyToId: utils.ParseStringToPointer(ln.InReplyToID, true),
// 		// Type: NULLABLE String (cast from an integer but not guaranteed to be a number) or null
// 		// Description: ID of the account that authored the status being replied to.
// 		InReplyToAccountId: utils.ParseStringToPointer(ln.InReplyToAccountId, true),
// 		// Type: NULLABLE Status or null
// 		// Description: The status being reblogged.
// 		Reblog: nil,
// 		// Type: NULLABLE Poll or null
// 		// Description: The poll attached to the status.
// 		Poll: nil,
// 		// Type: NULLABLE PreviewCard or null
// 		// Description: Preview card for links included within status content.
// 		Card: nil,
// 		// Type: NULLABLE String (ISO 639 Part 1 two-letter language code) or null
// 		// Description: Primary language of this status.
// 		Language: nil,
// 		// Type: NULLABLE String or null
// 		// Description: Plain-text source of a status. Returned instead of content when status is deleted, so the user may redraft from the source text without the client having to reverse-engineer the original text from the HTML content.
// 		Text: utils.ParseStringToPointer(ln.Status, true),
// 		// Type: NULLABLE String (ISO 8601 Datetime)
// 		// Description: Timestamp of when the status was last edited.
// 		EditedAt: nil,
// 		// Type: Boolean
// 		// Description: If the current token has an authorized user: Have you favourited this status?
// 		Favourited: false,
// 		// Type: Boolean
// 		// Description: If the current token has an authorized user: Have you boosted this status?
// 		Reblogged: false,
// 		// Type: Boolean
// 		// Description: If the current token has an authorized user: Have you muted notifications for this status’s conversation?
// 		Muted: false,
// 		// Type: Boolean
// 		// Description: If the current token has an authorized user: Have you bookmarked this status?
// 		Bookmarked: false,
// 		// Type: Boolean
// 		// Description: If the current token has an authorized user: Have you pinned this status? Only appears if the status is pinnable.
// 		Pinned: false,
// 		// Type: Array of FilterResult
// 		// Description: If the current token has an authorized user: The filter and keywords that matched this status.
// 		Filtered: nil,
// 	}
// 	return status
// }

// 好像不该在这里。
func ToMastodonRelationship(id, actor string) *entities.Relationship {
	lu := &core.LocalUser{
		ActivitypubID: id,
	}
	if err := dao.Read(lu); err != nil {
		log.Printf("%s", err.Error())
		return nil
	}
	acct := &entities.Account{
		Uri: id,
	}
	if err := dao.Read(acct); err != nil {
		log.Printf("%s", err.Error())
		return nil
	}
	// actor to object
	lra2o := &core.LocalRelation{
		Actor:  actor,
		Object: id,
	}
	if err := dao.Read(lra2o); err != nil {
		log.Printf("%s", err.Error())
		return nil
	}
	// object to actor
	lro2a := &core.LocalRelation{
		Actor:  id,
		Object: actor,
	}
	if err := dao.Read(lro2a); err != nil {
		log.Printf("%s", err.Error())
		return nil
	}

	relationship := &entities.Relationship{
		// Type: String (cast from an integer, but not guaranteed to be a number)
		// Description: The account ID.
		Attributes: utils.TimestampToRFC3339(lu.CreatedAt),
		// Type: Boolean
		// Description: Are you following this user?
		Following: lra2o.Type == core.RelationTypeFollow && lra2o.Status == core.RelationStatusAccepted,
		// Type: Boolean
		// Description: Are you receiving this user’s boosts in your home timeline?
		ShowingReblogs: true,
		// Type: Boolean
		// Description: Have you enabled notifications for this user?
		Notifying: false,
		// Type: Array of String (ISO 639-1 language two-letter code)
		// Description: Which languages are you following from this user?
		Languages: []string{},
		// Type: Boolean
		// Description: Are you followed by this user?
		FollowedBy: lro2a.Type == core.RelationTypeFollow && lro2a.Status == core.RelationStatusAccepted,
		// Type: Boolean
		// Description: Are you blocking this user?
		Blocking: lra2o.Type == core.RelationTypeBlock,
		// Type: Boolean
		// Description: Is this user blocking you?
		BlockedBy: lro2a.Type == core.RelationTypeBlock,
		// Type: Boolean
		// Description: Are you muting this user?
		Muting: false,
		// Type: Boolean
		// Description: Are you muting notifications from this user?
		MutingNotifications: false,
		// Type: Boolean
		// Description: Do you have a pending follow request for this user?
		Requested: lra2o.Type == core.RelationTypeFollow && lra2o.Status == core.RelationStatusPadding,
		// Type: Boolean
		// Description: Has this user requested to follow you?
		RequestedBy: lro2a.Type == core.RelationTypeFollow && lro2a.Status == core.RelationStatusPadding,
		// Type: Boolean
		// Description: Are you blocking this user’s domain?
		DomainBlocking: false,
		// Type: Boolean
		// Description: Are you featuring this user on your profile?
		Endorsed: false,
		// Type: String
		// Description: This user’s profile bio
		Note: acct.Note,
	}
	return relationship
}
