package activitypub

import "github.com/iancoleman/orderedmap"

func Follow(id, actor, object string) *orderedmap.OrderedMap {
	o := CreateOrderedMap([]*KV{
		{"@context", "https://www.w3.org/ns/activitystreams"},
		{"id", id},
		{"type", "Follow"},
		{"actor", actor},
		{"object", object},
	})
	return o
}
