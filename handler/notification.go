package handler

import (
	"github.com/Hana-ame/fedi-antenna/Tools/orderedmap"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

func Like(o *orderedmap.OrderedMap) error {
	name, host, _ := utils.ParseStatusesUriToNameHostTimestamp(o.GetOrDefault("object", "").(string))
	ln := &model.LocalNotify{
		ID:     o.GetOrDefault("id", "").(string),
		Actor:  o.GetOrDefault("actor", "").(string),
		Object: o.GetOrDefault("object", "").(string),
		Type:   model.NotifyTypeLike,
		To:     utils.ParseActivitypubID(name, host),
	}
	err := dao.Create(ln)
	return err
}

func Announce(o *orderedmap.OrderedMap) error {
	to := o.GetOrDefault("to", []string{}).([]string)
	cc := o.GetOrDefault("cc", []string{}).([]string)
	ln := &model.LocalNotify{
		ID:         o.GetOrDefault("id", "").(string),
		Actor:      o.GetOrDefault("actor", "").(string),
		Object:     o.GetOrDefault("object", "").(string),
		Type:       model.NotifyTypeAnnounce,
		Visibility: utils.ParseVisibility(to, cc),
		To:         utils.ParseTheOnlyUserFromToAndCc(to, cc),
	}
	err := dao.Create(ln)
	return err
}

func UndoNotify(o *orderedmap.OrderedMap) error {
	ln := &model.LocalNotify{
		ID:       o.GetOrDefault("id", "").(string),
		Actor:    o.GetOrDefault("actor", "").(string),
		Object:   o.GetOrDefault("object", "").(string),
		Type:     o.GetOrDefault("type", "").(string),
		DeleteAt: utils.Now(),
	}
	err := dao.Update(ln)
	return err
}
