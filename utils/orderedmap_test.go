package utils

import (
	"fmt"
	"testing"
)

func TestCreateOrderedMap(t *testing.T) {
	o := CreateOrderedMap([]*KV{
		{"1212", 123},
		{"123", 123},
	})
	fmt.Println(o)
}
