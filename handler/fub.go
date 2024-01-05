package handler

import (
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/orderedmap"
)

func Follow(o *orderedmap.OrderedMap) error {
	err := dao.Create(o)
	return err
}

func Undo(o *orderedmap.OrderedMap) error {
	err := dao.Delete(o)
	return err
}

func Block(o *orderedmap.OrderedMap) error {
	err := dao.Create(o)
	return err
}
