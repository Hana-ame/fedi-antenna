package activitypub

import "github.com/iancoleman/orderedmap"

func Undo(id, actor string, object *orderedmap.OrderedMap) *orderedmap.OrderedMap {
	o := CreateOrderedMap([]*KV{
		{"@context", "https://www.w3.org/ns/activitystreams"},
		{"id", id},
		{"type", "Undo"},
		{"actor", actor},
		{"object", object},
	})
	return o
}
