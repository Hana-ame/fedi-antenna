package fetch

import (
	"encoding/json"
	"net/http"

	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
)

// fetch remote user by ID
// ID is a valid url
func FetchUserByID(id string) (user *activitypub.User, err error) {
	resp, err := Fetch(http.MethodGet, id, nil, nil)
	if err != nil {
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return
	}

	user.IconURL = user.Icon.URL
	// log.Printf("%+v\n", user)
	return
}
