package entities

type PreviewCard struct {
	// Type: String (URL)
	// Description: Location of linked resource.
	Url string `json:"url"`
	// Type: String
	// Description: Title of linked resource.
	Title string `json:"title"`
	// Type: String
	// Description: Description of preview.
	Description string `json:"description"`
	// Type: String (Enumerable, oneOf)
	// Description: The type of the preview card.
	Type string `json:"type"`
	// Type: String
	// Description: The author of the original resource.
	AuthorName string `json:"author_name"`
	// Type: String (URL)
	// Description: A link to the author of the original resource.
	AuthorUrl string `json:"author_url"`
	// Type: String
	// Description: The provider of the original resource.
	ProviderName string `json:"provider_name"`
	// Type: String (URL)
	// Description: A link to the provider of the original resource.
	ProviderUrl string `json:"provider_url"`
	// Type: String (HTML)
	// Description: HTML to be used for generating the preview card.
	Html string `json:"html"`
	// Type: Integer
	// Description: Width of preview, in pixels.
	Width int `json:"width"`
	// Type: Integer
	// Description: Height of preview, in pixels.
	Height int `json:"height"`
	// Type: NULLABLE String (URL)
	// Description: Preview thumbnail.
	Image *string `json:"image"`
	// Type: String (URL)
	// Description: Used for photo embeds, instead of custom html.
	EmbedUrl string `json:"embed_url"`
	// Type: NULLABLE String
	// Description: A hash computed by the BlurHash algorithm, for generating colorful preview thumbnails when media has not been downloaded yet.
	Blurhash *string `json:"blurhash"`
}
