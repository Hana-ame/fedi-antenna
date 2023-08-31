package activitypub

import "github.com/iancoleman/orderedmap"

func Accept(id, actor string, object *orderedmap.OrderedMap) *orderedmap.OrderedMap {
	o := CreateOrderedMap([]*KV{
		{"@context", "https://www.w3.org/ns/activitystreams"},
		{"id", id},
		{"type", "Accept"},
		{"actor", actor},
		{"object", object},
	})
	return o
}
