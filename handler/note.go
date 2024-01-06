package handler

import (
	tools "github.com/Hana-ame/fedi-antenna/Tools"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/orderedmap"
)

func CreateNote(o *orderedmap.OrderedMap) error {
	n := &model.LocalNote{
		ID:           o.GetOrDefault("id", "").(string),
		AttributedTo: o.GetOrDefault("attributedTo", "").(string),
		Status:       o.GetOrDefault("content", "").(string),
		// todo
		MediaIDs:    nil,
		InReplyToID: tools.ParsePointerToString(o.Get("inReplyTo")),
		Sensitive:   o.GetOrDefault("sensitive", false).(bool),
		SpoilerText: tools.ParsePointerToString(o.Get("summary")),
		Visibility: utils.ParseVisibility(
			tools.ParseSliceToStringSlice(o.Get("to")),
			tools.ParseSliceToStringSlice(o.Get("cc")),
		),
		Published: (utils.Now()),
	}
	err := dao.Create(n)
	return err
}

func DeleteNote(id string) error {
	n := &model.LocalNote{
		ID:           id,
	}
	err := dao.Delete(n)
	return err
}
