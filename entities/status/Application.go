package status

type Application struct {
	// Type: String
	// Description: The name of the application that posted this status.
	Name string `json:"name"`
	// Type: NULLABLE String (URL) or null
	// Description: The website associated with the application that posted this status.
	Website *string `json:"website"`
}
