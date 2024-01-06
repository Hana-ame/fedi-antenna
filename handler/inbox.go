package handler

import (
	"fmt"
	"log"

	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/orderedmap"
)

// handlers

// public interface
// which o is the input object,
// user, host is the user
// err is the result of httpsig
func Inbox(o *orderedmap.OrderedMap, user, host string, err error) error {
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

func Create(o *orderedmap.OrderedMap) error {
	oo, ok := o.GetOrDefault("object", *orderedmap.New()).(orderedmap.OrderedMap)
	if !ok {
		log.Printf("create object do not have attribute object : %+v\n", o)
		return fmt.Errorf("create object do not have attribute object : %+v", o)
	}
	v, ok := oo.Get("type")
	if !ok {
		log.Printf("create object do not have type : %+v\n", o)
		return fmt.Errorf("create object do not have type : %+v", o)
	}
	s := v.(string)
	switch s {
	case activitypub.TypeNote:
		return CreateNote(&oo)
	default:
		log.Printf("create object have unknown type : %+v\n", s)
		return fmt.Errorf("create object have unknown type : %+v", s)
	}
}
