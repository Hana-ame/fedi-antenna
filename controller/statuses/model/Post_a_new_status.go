package model

// form data parameters
type Post_a_new_status struct {
	// REQUIRED String. The text content of the status. If media_ids is provided, this becomes optional. Attaching a poll is optional while status is provided.
	Status string `json:"status"`
	// REQUIRED Array of String. Include Attachment IDs to be attached as media. If provided, status becomes optional, and poll cannot be used.
	MediaIds []string `json:"media_ids"`

	Poll Poll `json:"poll"`
	// String. ID of the status being replied to, if status is a reply.
	InReplyToId *string `json:"in_reply_to_id"`
	// Boolean. Mark status and attached media as sensitive? Defaults to false.
	Sensitive bool `json:"sensitive"`
	// String. Text to be shown as a warning or subject before the actual content. Statuses are generally collapsed behind this field.
	SpoilerText string `json:"spoiler_text"`
	// String. Sets the visibility of the posted status to public, unlisted, private, direct.
	Visibility string `json:"visibility"`
	// String. ISO 639 language code for this status.
	Language string `json:"language"`
	// String. ISO 8601 Datetime at which to schedule a status. Providing this parameter will cause ScheduledStatus to be returned instead of Status. Must be at least 5 minutes in the future.
	ScheduledAt string `json:"scheduled_at"`
}

type Poll struct {
	// REQUIRED Array of String. Possible answers to the poll. If provided, media_ids cannot be used, and poll[expires_in] must be provided.
	Options []string `json:"options"`
	// REQUIRED Integer. Duration that the poll should be open, in seconds. If provided, media_ids cannot be used, and poll[options] must be provided.
	ExpiresIn int `json:"expires_in"`
	// Boolean. Allow multiple choices? Defaults to false.
	Multiple bool `json:"multiple"`
	// Boolean. Hide vote counts until the poll ends? Defaults to false.
	HideTotals bool `json:"hide_totals"`
}
