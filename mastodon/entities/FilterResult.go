package entities

type FilterResult struct {
	// Type: Filter
	// Description: The filter that was matched.
	Filter Filter `json:"filter"`
	// Type: NULLABLE Array of String, or null
	// Description: The keyword within the filter that was matched.
	KeywordMatches *[]string `json:"keyword_matches"`
	// Type: NULLABLE Array of String, or null
	// Description: The status ID within the filter that was matched.
	StatusMatches *[]string `json:"status_matches"`
}
