package activitypub

import "github.com/iancoleman/orderedmap"

func Tags(
	user, host string,
	items []any,
) *orderedmap.OrderedMap {
	o := CreateOrderedMap([]*KV{
		{"@context", "https://www.w3.org/ns/activitystreams"},
		{"id", APUserID(user, host) + "/collections/tags"},
		{"type", "Collection"},
		{"totalItems", len(items)},
		{"items", items},
	})
	return o
}
