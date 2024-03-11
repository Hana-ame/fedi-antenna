package handler

func inboxOfAccount(acitivitypubID string) (inbox string) {
	return acitivitypubID + "/inbox"
}
