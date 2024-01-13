package entities

type MediaAttachment struct {
	// Type: String (cast from an integer but not guaranteed to be a number)
	// Description: The ID of the attachment in the database.
	Id string `json:"id"`
	// Type: String (Enumerable, oneOf)
	// Description: The type of the attachment.
	Type string `json:"type"`
	// Type: String (URL)
	// Description: The location of the original full-size attachment.
	Url string `json:"url"`
	// Type: String (URL)
	// Description: The location of a scaled-down preview of the attachment.
	PreviewUrl string `json:"preview_url"`
	// Type: NULLABLE String (URL), or null if the attachment is local
	// Description: The location of the full-size original attachment on the remote website.
	RemoteUrl *string `json:"remote_url"`
	// Type: Hash
	// Description: Metadata returned by Paperclip.
	Meta map[string]any `json:"meta"`
	// Type: NULLABLE String, or null if alternate text was not provided for the media attachment
	// Description: Alternate text that describes what is in the media attachment, to be used for the visually impaired or when media attachments do not load.
	Description *string `json:"description"`
	// Type: String (Blurhash)
	// Description: A hash computed by the BlurHash algorithm, for generating colorful preview thumbnails when media has not been downloaded yet.
	Blurhash string `json:"blurhash"`
}
