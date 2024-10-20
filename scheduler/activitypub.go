package scheduler

import (
	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
	"github.com/Hana-ame/fedi-antenna/activitypub/curl"
)

func ByApID(id string) (*orderedmap.OrderedMap, error) {
	return curl.GetWithSignDefault(id)
}
