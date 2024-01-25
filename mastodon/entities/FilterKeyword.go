package entities

type FilterKeyword struct {
	// Type: String (cast from an integer, but not guaranteed to be a number)
	// Description: The ID of the FilterKeyword in the database.
	Id string `json:"id"`
	// Type: String
	// Description: The phrase to be matched against.
	Keyword string `json:"keyword"`
	// Type: Boolean
	// Description: Should the filter consider word boundaries? See implementation guidelines for filters.
	WholeWord bool `json:"whole_word"`
}
