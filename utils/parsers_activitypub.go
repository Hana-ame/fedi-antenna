package utils

import (
	"strings"

	"github.com/google/uuid"
)

// "https://" + host + "/users/" + name
func ParseActivitypubID(name, host string) string {
	return "https://" + host + "/users/" + name
}

// "https://" + host + "/@" + name
func ParseProfileUrl(name, host string) string {
	return "https://" + host + "/@" + name
}

// return  host, name  found from activitypubID
func ParseNameAndHost(activitypubID string) (name, host string) {
	activitypubID = strings.TrimPrefix(activitypubID, "https://")
	strSlince := strings.Split(activitypubID, "/")
	if len(strSlince) > 0 {
		return strSlince[len(strSlince)-1], strSlince[0]
	}
	return "", ""
}

func GenerateObjectID(typ, host string) string {
	return "https://" + host + "/objects/" + typ + "/" + uuid.New().String()
}
