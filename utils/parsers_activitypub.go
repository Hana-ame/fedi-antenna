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

// [host, name] found from activitypubID
func ParseNameAndHost(activitypubID string) (name, host string) {
	activitypubID = strings.TrimPrefix(activitypubID, "https://")
	strSlince := strings.Split(activitypubID, "/")
	if len(strSlince) > 0 {
		return strSlince[len(strSlince)-1], strSlince[0]
	}
	return "", ""
}

// "https://" + host + "/objects/" + typ + "/" + uuid.New().String()
func GenerateObjectID(typ, host string) string {
	return "https://" + host + "/objects/" + typ + "/" + uuid.New().String()
}

// only for the query of local objects
func ParseTypeFromObjectID(id string) (typ, host string) {
	id = strings.TrimPrefix(id, "https://")
	arr := strings.Split(id, "/")
	if len(arr) < 3 {
		return
	}
	return arr[2], arr[0]
}

// "https://" + host + "/users/" + name + "/statues/" + id
func ParseStatusesID(name, host, id string) ( statusesID string) {
	return "https://" + host + "/users/" + name + "/statues/" + id
}

// "https://" + host + "/users/" + name + "/statues/" + timestamp
func ParseStatusesNameHostTimestamp(statusesID string) (name, host, timestamp string) {
	statusesID = strings.TrimPrefix(statusesID, "https://")
	arr := strings.Split(statusesID, "/")
	if len(arr) < 5 {
		return
	}
	return arr[2], arr[0], arr[4]
}

// "https://" + host + "/@" + name + "/" + timestamp
func ParseStatusesURL(name, host, timestamp string) (statusesID string ){
	return "https://" + host + "/@" + name + "/" + timestamp
}
