package handler

import (
	"fmt"
	"log"

	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
)

func Follow(o *orderedmap.OrderedMap) error {
	r := &model.LocalRelation{
		ID:     o.GetOrDefault("id", "").(string),
		Actor:  o.GetOrDefault("actor", "").(string),
		Object: o.GetOrDefault("object", "").(string),
		Type:   model.RelationTypeFollow,
		Status: model.RelationStatusPadding,
	}
	err := dao.Create(r)
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
	r := &model.LocalRelation{
		ID:     oo.GetOrDefault("id", "").(string),
		Actor:  o.GetOrDefault("actor", "").(string),
		Object: o.GetOrDefault("object", "").(string),
		Type:   model.RelationTypeFollow,
		Status: model.RelationStatusAccepted,
	}
	err := dao.Update(r)
	return err
}

func Reject(o *orderedmap.OrderedMap) error {
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
	r := &model.LocalRelation{
		ID:     oo.GetOrDefault("id", "").(string),
		Actor:  o.GetOrDefault("actor", "").(string),
		Object: o.GetOrDefault("object", "").(string),
		Type:   model.RelationTypeFollow,
		Status: model.RelationStatusRejected,
	}
	err := dao.Update(r)
	return err
}

func Block(o *orderedmap.OrderedMap) error {
	r := &model.LocalRelation{
		ID:     o.GetOrDefault("id", "").(string),
		Actor:  o.GetOrDefault("actor", "").(string),
		Object: o.GetOrDefault("object", "").(string),
		Type:   model.RelationTypeBlock,
		Status: model.RelationStatusPadding,
	}
	err := dao.Create(r)
	return err
}

func UndoRelation(o *orderedmap.OrderedMap) error {
	r := &model.LocalRelation{
		ID:     o.GetOrDefault("id", "").(string),
		Actor:  o.GetOrDefault("actor", "").(string),
		Object: o.GetOrDefault("object", "").(string),
		Type:   o.GetOrDefault("type", "").(string),
		Status: model.RelationStatusUndo,
	}
	err := dao.Update(r)
	return err
}
