package activitypub

import (
	"github.com/Hana-ame/fedi-antenna/utils"
	"github.com/iancoleman/orderedmap"
)

func Emoji(id, name string, updated int64, icon *orderedmap.OrderedMap) *orderedmap.OrderedMap {
	o := CreateOrderedMap([]*KV{
		{"id", id},
		{"type", "Emoji"},
		{"name", name},
		{"updated", utils.TimestampToRFC3339(updated)},
		{"icon", icon},
	})
	return o
}
