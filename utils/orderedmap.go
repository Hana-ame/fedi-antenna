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

// utils
func OrderedMap(keys []string, values []interface{}) *orderedmap.OrderedMap {
	o := orderedmap.New()
	for i, key := range keys {
		o.Set(key, values[i])
	}
	return o
}

func IdType(id, typestr string) *orderedmap.OrderedMap {
	return OrderedMap([]string{"@id", "@type"}, []any{id, typestr})
}

// mycrypto
func PkPem(name string) string {
	pk, _ := ReadKeyFromFile(name + ".pem")
	pubKeyStr := MarshalPublicKey(&pk.PublicKey)
	return string(pubKeyStr)
}
