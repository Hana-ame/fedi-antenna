package activitypub

import (
	"github.com/Hana-ame/fedi-antenna/utils"
	"github.com/iancoleman/orderedmap"
)

// icon

func ImageObj(mediaType, url string) *orderedmap.OrderedMap {
	o := utils.OrderedMap([]string{
		"type",
		"mediaType",
		"url",
	}, []any{
		"Image",
		mediaType,
		url,
	})
	return o
}
