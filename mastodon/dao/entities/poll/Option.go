package poll

type Option struct {
	// Type: String
	// Description: The text value of the poll option.
	Title string `json:"title"`
	// Type: NULLABLE Integer, or null if results are not published yet.
	// Description: The total number of received votes for this option.
	VotesCount *int `json:"votes_count"`
}
