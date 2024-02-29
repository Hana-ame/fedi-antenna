package handler

import (
	"fmt"
	"log"

	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
)

func Follow(o *orderedmap.OrderedMap) error {
	err := dao.Follow(
		o.GetOrDefault("id", "").(string),
		o.GetOrDefault("object", "").(string),
		o.GetOrDefault("actor", "").(string),
	)
	return err
}

func Accept(o *orderedmap.OrderedMap) error {
	oo, ok := o.GetOrDefault("object", *orderedmap.New()).(orderedmap.OrderedMap)
	if !ok {
		log.Printf("not exist attribute object in Accept object: %+v\n", o)
		return fmt.Errorf("not exist attribute object in Accept object: %+v", o)
	}
	s, ok := oo.GetOrDefault("type", "unknown").(string)
	if !ok || s != "Follow" {
		log.Printf("accept object do not have type : %+v\n", o)
		return fmt.Errorf("accept object do not have type : %+v", o)
	}
	err := dao.Accept(
		oo.GetOrDefault("id", "").(string),
		o.GetOrDefault("object", "").(string),
		o.GetOrDefault("actor", "").(string),
	)
	return err
}

func Reject(o *orderedmap.OrderedMap) error {
	oo, ok := o.GetOrDefault("object", *orderedmap.New()).(orderedmap.OrderedMap)
	if !ok {
		log.Printf("not exist attribute object in Accept object: %+v\n", o)
		return fmt.Errorf("not exist attribute object in Accept object: %+v", o)
	}
	s, ok := oo.GetOrDefault("type", "").(string)
	if !ok || s != "Follow" {
		log.Printf("accept object do not have type : %+v\n", o)
		return fmt.Errorf("accept object do not have type : %+v", o)
	}
	err := dao.Reject(
		oo.GetOrDefault("id", "").(string),
		o.GetOrDefault("object", "").(string),
		o.GetOrDefault("actor", "").(string),
	)
	return err
}

func Block(o *orderedmap.OrderedMap) error {
	err := dao.Block(
		o.GetOrDefault("id", "").(string),
		o.GetOrDefault("object", "").(string),
		o.GetOrDefault("actor", "").(string))
	return err
}

func UndoRelation(o *orderedmap.OrderedMap) error {
	switch o.GetOrDefault("type", "").(string) {
	case model.RelationTypeBlock:
		err := dao.Unblock(
			o.GetOrDefault("id", "").(string),
			o.GetOrDefault("object", "").(string),
			o.GetOrDefault("actor", "").(string))
		return err
	case model.RelationTypeFollow:
		err := dao.Unfollow(
			o.GetOrDefault("id", "").(string),
			o.GetOrDefault("object", "").(string),
			o.GetOrDefault("actor", "").(string))
		return err
	default:
		return nil
	}
	// err := dao.Update(r)
	// return nil
}
