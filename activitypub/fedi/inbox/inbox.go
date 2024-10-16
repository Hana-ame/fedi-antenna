package inbox

import (
	"fmt"
	"log"

	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
	"github.com/Hana-ame/fedi-antenna/activitypub/dao/model"
	"github.com/Hana-ame/fedi-antenna/utils"
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
	case model.TypeCreate:
		return Create(o)
	case model.TypeDelete:
		return Delete(o)
	case model.TypeFollow:
		return Follow(o)
	case model.TypeBlock:
		return Block(o)
	case model.TypeUndo:
		return Undo(o)
	case model.TypeReject:
		return Reject(o)
	case model.TypeAccept:
		return Accept(o)
	case model.TypeLike:
		return Like(o)
	case model.TypeAnnounce:
		return Announce(o)
	default:
		log.Printf("inbox have unknown type : %+v\n", s)
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
	case model.TypeNote:
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
	case model.TypeBlock, model.TypeFollow:
		return UndoRelation(&oo)
	case model.TypeAnnounce, model.TypeLike:
		return UndoNotify(&oo)
	default:
		log.Printf("undo object have unknown type : %+v\n", o)
		return fmt.Errorf("undo object have unknown type : %+v", o)
	}
}

func Delete(o *orderedmap.OrderedMap) error {
	id, typ := utils.ParseObjectIDAndType(o)
	switch typ {
	case model.TypeNote:
		return DeleteNote(id)
	case model.TypePerson:
		return DeletePerson(id)
	default:
		log.Printf("delete object have unknown type : %+v\n", o)
		return fmt.Errorf("delete object have unknown type : %+v", o)
	}
}
