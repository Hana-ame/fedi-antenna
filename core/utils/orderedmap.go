package utils

import "github.com/Hana-ame/fedi-antenna/Tools/orderedmap"

type KV struct {
	Key   string
	Value any
}

func NewMapFromKV(list []*KV) *orderedmap.OrderedMap {
	o := orderedmap.New()
	for _, p := range list {
		o.Set(p.Key, p.Value)
	}
	return o
}
