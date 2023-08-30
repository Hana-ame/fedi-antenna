package activitypub

import "github.com/iancoleman/orderedmap"

// id is object
func Block(id, actor, object string) *orderedmap.OrderedMap {
	o := CreateOrderedMap([]*KV{
		{"@context", "https://www.w3.org/ns/activitystreams"},
		{"id", id},
		{"type", "Block"},
		{"actor", actor},
		{"object", object},
	})
	return o
}
