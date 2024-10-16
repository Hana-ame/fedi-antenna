package account

type Field struct {
	// Type: String
	// Description: The key of a given field’s key-value pair.
	Name string `json:"name"`
	// Type: String (HTML)
	// Description: The value associated with the name key.
	Value string `json:"value"`
	// Type: NULLABLE String (ISO 8601 Datetime) if value is a verified URL. Otherwise, null.
	// Description: Timestamp of when the server verified a URL value for a rel=“me” link.
	VerifiedAt *string `json:"verified_at"`
}
