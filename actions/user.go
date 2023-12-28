package actions

import (
	"encoding/json"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/activitypub/model"
)

// fetch remote user by ID
// ID is a valid url
func FetchUserByID(id string) (user *model.User, err error) {
	resp, err := Fetch(http.MethodGet, id, nil, nil)
	if err != nil {
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return
	}

	return
}
