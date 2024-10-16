package entities

type Filter struct {
	// Type: String (cast from an integer, but not guaranteed to be a number)
	// Description: The ID of the Filter in the database.
	Id string `json:"id"`
	// Type: String
	// Description: A title given by the user to name the filter.
	Title string `json:"title"`
	// Type: Array of String (Enumerable, anyOf)
	// Description: The contexts in which the filter should be applied.
	Context []string `json:"context"`
	// Type: NULLABLE String (ISO 8601 Datetime), or null if the filter does not expire
	// Description: When the filter should no longer be applied.
	ExpiresAt *string `json:"expires_at"`
	// Type: String (Enumerable, oneOf)
	// Description: The action to be taken when a status matches this filter.
	FilterAction string `json:"filter_action"`
	// Type: Array of FilterKeyword
	// Description: The keywords grouped under this filter.
	Keywords []FilterKeyword `json:"keywords"`
	// Type: Array of FilterStatus
	// Description: The statuses grouped under this filter.
	Statuses []FilterStatus `json:"statuses"`
}
