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
