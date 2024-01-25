package utils

import (
	"strings"

	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
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
// "https://" + host + "/users/" + name + "/xxx"
func ParseNameAndHost(activitypubID string) (name, host string) {
	activitypubID = strings.TrimPrefix(activitypubID, "https://")
	strSlince := strings.Split(activitypubID, "/")
	if len(strSlince) < 1 {
		return
	}
	if (len(strSlince)) < 3 {
		return "", strSlince[0]
	}
	return strSlince[2], strSlince[0]
}

// "https://" + host + "/objects/" + typ + "/" + uuid.New().String()
func GenerateObjectID(typ, host string) string {
	return "https://" + host + "/objects/" + typ + "/" + uuid.New().String()
}

// only for the query of local objects
// https:// + host + /object/ + typ
func ParseTypeFromObjectID(id string) (typ, host string) {
	id = strings.TrimPrefix(id, "https://")
	arr := strings.Split(id, "/")
	if len(arr) < 3 {
		return
	}
	return arr[2], arr[0]
}

// "https://" + host + "/users/" + name + "/statuses/" + id
func ParseStatusesUri(name, host, id string) (statusesUri string) {
	return "https://" + host + "/users/" + name + "/statuses/" + id
}

// "https://" + host + "/users/" + name + "/statuses/" + timestamp (+ "/xxx")
func ParseStatusesUriToNameHostTimestamp(statusesUri string) (name, host, timestamp string) {
	statusesUri = strings.TrimPrefix(statusesUri, "https://")
	arr := strings.Split(statusesUri, "/")
	if len(arr) < 3 {
		return
	} else if (len(arr)) < 5 {
		return arr[2], arr[0], ""
	}
	return arr[2], arr[0], arr[4]
}

// "https://" + host + "/@" + name + "/" + timestamp
func ParseStatusesURL(name, host, timestamp string) (statusesID string) {
	return "https://" + host + "/@" + name + "/" + timestamp
}

// kari
func ParseVisibility(to, cc []string) string {
	publicStream := "https://www.w3.org/ns/activitystreams#Public"
	for _, v := range to {
		if v == publicStream {
			return "public"
		}
	}
	for _, v := range cc {
		if v == publicStream {
			return "unlisted"
		}
	}
	for _, v := range to {
		if strings.HasSuffix(v, "follower") || strings.HasSuffix(v, "followers") {
			return "private"
		}
	}
	return "direct"
}

// kari
func ParseObjectIDAndType(o *orderedmap.OrderedMap) (id, typ string) {
	if object, ok := o.Get("object"); ok {
		if s, ok := object.(string); ok {
			return s, "Person"
		} else if o, ok := object.(orderedmap.OrderedMap); ok {
			if id, ok := o.Get("id"); ok {
				if s, ok := id.(string); ok {
					return s, "Note"
				}
			}
		}
	}
	return
}

// kari
func ParseTheOnlyUserFromToAndCc(to []string, cc []string) string {
	var userID string
	for _, s := range to {
		if strings.Contains(s, "/users/") {
			if userID == "" {
				userID = s
			} else {
				return ""
			}
		}
	}
	for _, s := range cc {
		if strings.Contains(s, "/users/") {
			if userID == "" {
				userID = s
			} else {
				return ""
			}
		}
	}
	return userID
}
