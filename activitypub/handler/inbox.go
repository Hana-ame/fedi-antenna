package handler

import (
	"fmt"
	"log"

	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

// handlers

// public interface
// which o is the input object,
// user is the user that inbox's owner.
// host is the host that local server on.
// err is the result of httpsig, <nil> for pass
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
	case activitypub.TypeDelete:
		return Delete(o)
	case activitypub.TypeFollow:
		return Follow(o)
	case activitypub.TypeBlock:
		return Block(o)
	case activitypub.TypeUndo:
		return Undo(o)
	case activitypub.TypeReject:
		return Reject(o)
	case activitypub.TypeAccept:
		return Accept(o)
	case activitypub.TypeLike:
		return Like(o)
	case activitypub.TypeAnnounce:
		return Announce(o)
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

func Undo(o *orderedmap.OrderedMap) error {
	oo, ok := o.GetOrDefault("object", *orderedmap.New()).(orderedmap.OrderedMap)
	if !ok {
		log.Printf("not exist attribute object in undo object: %+v\n", o)
		return fmt.Errorf("not exist attribute object in undo object: %+v", o)
	}
	s, ok := oo.GetOrDefault("type", "unknown").(string)
	if !ok {
		log.Printf("undo object do not have type : %+v\n", o)
		return fmt.Errorf("undo object do not have type : %+v", o)
	}
	switch s {
	case activitypub.TypeBlock, activitypub.TypeFollow:
		return UndoRelation(&oo)
	case activitypub.TypeAnnounce, activitypub.TypeLike:
		return UndoNotify(&oo)
	default:
		log.Printf("undo object have unknown type : %+v\n", o)
		return fmt.Errorf("undo object have unknown type : %+v", o)
	}
}

func Delete(o *orderedmap.OrderedMap) error {
	id, typ := utils.ParseObjectIDAndType(o)
	switch typ {
	case activitypub.TypeNote:
		return DeleteNote(id)
	case activitypub.TypePerson:
		return DeletePerson(id)
	default:
		log.Printf("delete object have unknown type : %+v\n", o)
		return fmt.Errorf("delete object have unknown type : %+v", o)
	}
}
