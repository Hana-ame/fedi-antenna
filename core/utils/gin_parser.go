package utils

import (
	"strconv"
)

// return &s
func StringToPointer(s string, emptyIsNil bool) *string {
	if emptyIsNil && s == "" {
		return nil
	}
	return &s
}

// return &b
func BoolToPointer(b bool, emptyIsNil bool) *bool {
	if emptyIsNil && !b {
		return nil
	}
	return &b
}

func PointerToBool(pb *bool, def bool) bool {
	if pb == nil {
		return def
	}
	return *pb
}
func PointerToString(ps *string, def string) string {
	if ps == nil {
		return def
	}
	return *ps
}

func Atoi(s string, def int) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		return def
	}
	return i
}
