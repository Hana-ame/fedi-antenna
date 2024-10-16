package inbox

import (
	"fmt"
	"log"

	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
	"github.com/Hana-ame/fedi-antenna/activitypub/dao"
	"github.com/Hana-ame/fedi-antenna/activitypub/dao/model"
	"github.com/Hana-ame/fedi-antenna/utils"
)

func Like(o *orderedmap.OrderedMap) error {
	err := dao.Like(o.GetOrDefault("id", "").(string), o.GetOrDefault("actor", "").(string), o.GetOrDefault("object", "").(string), utils.Now())
	return err
}

func Announce(o *orderedmap.OrderedMap) error {
	to := o.GetOrDefault("to", []string{}).([]string)
	cc := o.GetOrDefault("cc", []string{}).([]string)
	visibility := utils.ParseVisibility(to, cc)
	err := dao.Announce(o.GetOrDefault("id", "").(string), o.GetOrDefault("actor", "").(string), o.GetOrDefault("object", "").(string), visibility, utils.Now())
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
	case model.TypeLike:
		err := dao.UndoLike(o.GetOrDefault("id", "").(string), o.GetOrDefault("actor", "").(string), o.GetOrDefault("object", "").(string), utils.Now())
		return err
	case model.TypeAnnounce:
		err := dao.UndoAnnounce(o.GetOrDefault("id", "").(string), o.GetOrDefault("actor", "").(string), o.GetOrDefault("object", "").(string), utils.Now())
		return err
	default:
		log.Printf("not support, %+v", o)
		return fmt.Errorf("not support, %+v", o)
	}
}
