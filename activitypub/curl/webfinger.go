package curl

import (
	"fmt"
	"net/http"
	"strings"

	tools "github.com/Hana-ame/fedi-antenna/Tools"
	mycurl "github.com/Hana-ame/fedi-antenna/Tools/my_curl"
	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
)

// @nanakananoka@mstdn.jp | nanakananoka@mstdn.jp
func Webfinger(acct string) (*orderedmap.OrderedMap, error) {
	acct = strings.TrimPrefix(acct, "@")
	arr := strings.Split(acct, "@")
	if len(arr) < 2 {
		return nil, fmt.Errorf("the format of acct is incorrect:" + acct)
	}
	username, host := arr[0], arr[1]

	status, body, err := mycurl.Get("", nil, "", `https://`+host+`/.well-known/webfinger?resource=acct:`+username+`@`+host)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("return http status %d", status)
	}
	o, err := tools.BytesToJson(body)
	return o, err
}
