package activitypub

import "github.com/iancoleman/orderedmap"

func Reject(id, actor string, object *orderedmap.OrderedMap) *orderedmap.OrderedMap {
	o := CreateOrderedMap([]*KV{
		{"@context", "https://www.w3.org/ns/activitystreams"},
		{"id", id},
		{"type", "Reject"},
		{"actor", actor},
		{"object", object},
	})
	return o
}
