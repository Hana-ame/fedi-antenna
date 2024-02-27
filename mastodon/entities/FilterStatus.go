package entities

type FilterStatus struct {
	// Type: String (cast from an integer, but not guaranteed to be a number)
	// Description: The ID of the FilterStatus in the database.
	Id string `json:"id"`
	// Type: String (cast from an integer, but not guaranteed to be a number)
	// Description: The ID of the Status that will be filtered.
	StatusId string `json:"status_id"`
}
