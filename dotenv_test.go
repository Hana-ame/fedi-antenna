package main

import (
	"fmt"
	"testing"

	tools "github.com/Hana-ame/fedi-antenna/Tools"
)

func TestXxx(t *testing.T) {
	m := make(map[string]string)
	json, _ := tools.ReadJsonFromFile("hosts.json")
	hosts := InterfaceSliceToStringSlice(json.GetOrDefault("hosts", []any{}).([]any))
	fmt.Println(hosts)
	for _, k := range json.Keys() {
		if k == "hosts" {
			continue
		}
		v := json.GetOrDefault(k, k).(string)
		m[k] = v
	}
	fmt.Println(m)
}

func InterfaceSliceToStringSlice(input []interface{}) []string {
	var output []string

	for _, v := range input {
		if str, ok := v.(string); ok {
			output = append(output, str)
		}
	}

	return output
}
