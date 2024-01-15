package model

// form data parameters
type Follow_account struct {
	// Boolean. Receive this account’s reblogs in home timeline? Defaults to true.
	Reblogs bool `json:"reblogs"`
	// Boolean. Receive notifications when this account posts a status? Defaults to false.
	Notify bool `json:"notify"`
	// Array of String (ISO 639-1 language two-letter code). Filter received statuses for these languages. If not provided, you will receive this account’s posts in all languages.
	Languages []string `json:"languages"`
}
