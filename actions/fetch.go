package actions

import (
	"io"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/Tools/myfetch"
	"github.com/Hana-ame/fedi-antenna/core/utils"
) 

func fetch(method, url string, header map[string]string, body io.Reader) (*http.Response, error) {
	header = utils.MergeMaps(header, map[string]string{"Accept": "application/activity+json"})
	return myfetch.Fetch(method, url, header, body)
}
