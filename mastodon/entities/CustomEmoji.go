package entities

type CustomEmoji struct {
	// Type: String
	// Description: The name of the custom emoji.
	Shortcode string `json:"shortcode"`
	// Type: String (URL)
	// Description: A link to the custom emoji.
	Url string `json:"url"`
	// Type: String (URL)
	// Description: A link to a static copy of the custom emoji.
	StaticUrl string `json:"static_url"`
	// Type: Boolean
	// Description: Whether this Emoji should be visible in the picker or unlisted.
	VisibleInPicker bool `json:"visible_in_picker"`
	// Type: String
	// Description: Used for sorting custom emoji in the picker.
	Category string `json:"category"`
}
