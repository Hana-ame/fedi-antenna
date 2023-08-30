package utils

import "github.com/iancoleman/orderedmap"

type KV struct {
	Key   string
	Value any
}

func CreateOrderedMap(kvlist []*KV) (o *orderedmap.OrderedMap) {
	o = orderedmap.New()
	for _, v := range kvlist {
		o.Set(v.Key, v.Value)
	}
	return
}
