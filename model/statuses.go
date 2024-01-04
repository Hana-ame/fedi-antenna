package model

// not sure.
type Poll struct {
	Options    []string `json:"options"`
	ExpiresIn  int      `json:"expires_in"`
	Multiple   bool     `json:"multiple,omitempty"`
	HideTotals bool     `json:"hide_totals,omitempty"`
}
type Status struct {
	Status      string   `json:"status"`
	MediaIDs    []string `json:"media_ids"`
	Poll        *Poll    `json:"poll"`
	InReplyToID string   `json:"in_reply_to_id,omitempty"`
	Sensitive   bool     `json:"sensitive"`
	SpoilerText string   `json:"spoiler_text"`
	Visibility  string   `json:"visibility"`
	Language    string   `json:"language"`
	ScheduledAt string   `json:"scheduled_at"`
}
