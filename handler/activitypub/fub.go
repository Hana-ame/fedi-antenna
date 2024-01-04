package handler

import (
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
)

func Follow(o *activitypub.Follow) error {
	err := dao.Create(o)
	return err
}

func Undo(o *activitypub.Undo) error {
	err := dao.Delete(o.Object)
	return err
}

func Block(o *activitypub.Block) error {
	err := dao.Create(o)
	return err
}
