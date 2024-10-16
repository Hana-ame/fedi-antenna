package inbox

import (
	"fmt"
	"log"

	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
	"github.com/Hana-ame/fedi-antenna/activitypub/dao"
	"github.com/Hana-ame/fedi-antenna/activitypub/dao/model"
	"github.com/Hana-ame/fedi-antenna/utils"
)

func Follow(o *orderedmap.OrderedMap) error {
	err := dao.Follow(
		o.GetOrDefault("id", "").(string),
		o.GetOrDefault("object", "").(string),
		o.GetOrDefault("actor", "").(string),
		utils.Now(),
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
		oo.GetOrDefault("actor", "").(string),
		oo.GetOrDefault("object", "").(string),
		utils.Now(),
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
		oo.GetOrDefault("actor", "").(string),
		oo.GetOrDefault("object", "").(string),
		utils.Now(),
	)
	return err
}

func Block(o *orderedmap.OrderedMap) error {
	err := dao.Block(
		o.GetOrDefault("id", "").(string),
		o.GetOrDefault("actor", "").(string),
		o.GetOrDefault("object", "").(string),
		utils.Now(),
	)
	return err
}

func UndoRelation(o *orderedmap.OrderedMap) error {
	switch o.GetOrDefault("type", "").(string) {
	case model.TypeBlock:
		err := dao.UndoBlock(
			o.GetOrDefault("id", "").(string),
			o.GetOrDefault("actor", "").(string),
			o.GetOrDefault("object", "").(string),
			utils.Now(),
		)
		return err
	case model.TypeFollow:
		err := dao.UndoFollow(
			o.GetOrDefault("id", "").(string),
			o.GetOrDefault("actor", "").(string),
			o.GetOrDefault("object", "").(string),
			utils.Now(),
		)
		return err
	default:
		log.Printf("unknown type : %+v\n", o)
		return fmt.Errorf("unknown type : %+v", o)
	}
	// err := dao.Update(r)
	// return nil
}
