// utils
package s2s

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/activitypub"
	"github.com/Hana-ame/fedi-antenna/utils"
	"github.com/iancoleman/orderedmap"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func signObjectByUser(object *orderedmap.OrderedMap, user, host, endpoint string) (*http.Request, error) {
	body, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	pk, err := utils.ReadKeyFromFile(user + ".pem")
	if err != nil {
		return nil, err
	}
	req, err := utils.NewSingedRequest(pk, activitypub.APUserID(user, host)+"#main-key", http.MethodPost, endpoint, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func fetchObject(url string) (o *orderedmap.OrderedMap, err error) {
	return utils.FetchObj(http.MethodGet, url, nil)
}

// todo
func fetchObjectWithCache(url string) (o *orderedmap.OrderedMap, err error) {
	return fetchObject(url)
}

func fetchObjectWithRequest(r *http.Request) (o *orderedmap.OrderedMap, err error) {
	resp, err := utils.Do(r)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, nil
	}
	o = orderedmap.New()
	err = json.Unmarshal(body, &o)
	if err != nil {
		return nil, err
	}
	return
}
