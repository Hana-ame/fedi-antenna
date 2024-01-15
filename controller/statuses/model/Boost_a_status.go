package model

// form data parameters
type Boost_a_status struct {
	// String. Any visibility except limited or direct (i.e. public, unlisted, private). Defaults to public. Currently unused in UI.
	Visibility string `json:"visibility"`
}
