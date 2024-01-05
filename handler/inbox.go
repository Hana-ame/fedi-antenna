package handler

import (
	"fmt"
	"log"

	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/orderedmap"
)

// handlers

// public interface
func Inbox(o *orderedmap.OrderedMap, user string) error {
	v, ok := o.Get("type")
	if !ok {
		log.Printf("inbox do not have type : %+v", o)
		return fmt.Errorf("inbox do not have type : %+v", o)
	}
	s := v.(string)
	switch s {
	case activitypub.TypeCreate:
		return Create(o)
	case activitypub.TypeFollow:
		return Follow(o)
	case activitypub.TypeBlock:
		return Block(o)
	case activitypub.TypeUndo:
		return Undo(o)
	default:
		log.Printf("inbox have unknown type : %+v", s)
		return fmt.Errorf("inbox have unknown type : %+v", s)
	}

	// return nil
}
