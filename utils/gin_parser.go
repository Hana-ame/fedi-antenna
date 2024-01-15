package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Parser struct {
	*gin.Context
}

func (p *Parser) PostForm(key string, value *string) error {
	var ok bool
	if *value, ok = p.Context.GetPostForm(key); !ok {
		return fmt.Errorf("PostForm should have key: %s", key)
	}
	return nil
}

// return &s
func ParseStringToPointer(s string, emptyIsNil bool) *string {
	if emptyIsNil && s == "" {
		return nil
	}
	return &s
}

// return &s
func ParseBoolToPointer(b bool, emptyIsNil bool) *bool {
	if emptyIsNil && !b {
		return nil
	}
	return &b
}