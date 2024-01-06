package handler

import (
	"fmt"
	"log"

	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/orderedmap"
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

func Undo(o *orderedmap.OrderedMap) error {
	oo, ok := o.GetOrDefault("object", nil).(orderedmap.OrderedMap)
	if !ok {
		log.Printf("not exist attribute object in undo object: %+v\n", o)
		return fmt.Errorf("not exist attribute object in undo object: %+v", o)
	}

	r := &model.LocalRelation{
		ID: oo.GetOrDefault("id", "").(string),
	}
	err := dao.Delete(r)
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
