package model

// form data parameters
type Register_an_account struct {
	// REQUIRED String. The desired username for the account
	Username string `json:"username"`
	// REQUIRED String. The email address to be used for login
	Email string `json:"email"`
	// REQUIRED String. The password to be used for login
	Password string `json:"password"`
	// REQUIRED Boolean. Whether the user agrees to the local rules, terms, and policies. These should be presented to the user in order to allow them to consent before setting this parameter to TRUE.
	Agreement bool `json:"agreement"`
	// REQUIRED String. The language of the confirmation email that will be sent.
	Locale string `json:"locale"`
	// String. If registrations require manual approval, this text will be reviewed by moderators.
	Reason string `json:"reason"`
}
