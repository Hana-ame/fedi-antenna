package handler

import "log"

func inboxOfAccount(acitivitypubID string) (inbox string) {
	return acitivitypubID + "/inbox"
}

func logE(err error) {
	log.Println(err)
}
