package handler

import (
	"fmt"
	"log"

	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

func Like(o *orderedmap.OrderedMap) error {
	err := core.Favourite(o.GetOrDefault("id", "").(string), o.GetOrDefault("object", "").(string), o.GetOrDefault("actor", "").(string))
	return err
}

func Announce(o *orderedmap.OrderedMap) error {
	to := o.GetOrDefault("to", []string{}).([]string)
	cc := o.GetOrDefault("cc", []string{}).([]string)
	_, err := core.Reblog(o.GetOrDefault("id", "").(string), o.GetOrDefault("object", "").(string), o.GetOrDefault("actor", "").(string), utils.ParseVisibility(to, cc))
	return err
}

func UndoNotify(o *orderedmap.OrderedMap) error {
	// ln := &model.LocalNotify{
	// 	ID:       o.GetOrDefault("id", "").(string),
	// 	Actor:    o.GetOrDefault("actor", "").(string),
	// 	Object:   o.GetOrDefault("object", "").(string),
	// 	Type:
	// 	DeleteAt: utils.Now(),
	// }
	// err := dao.Update(ln)
	switch o.GetOrDefault("type", "").(string) {
	case model.NotifyTypeAnnounce:
		err := core.Unfavourite(o.GetOrDefault("object", "").(string), o.GetOrDefault("actor", "").(string))
		return err
	case model.NotifyTypeLike:
		err := core.Unreblog(o.GetOrDefault("object", "").(string), o.GetOrDefault("actor", "").(string))
		return err
	default:
		log.Println("not support")
		return fmt.Errorf("not support")
	}
}
