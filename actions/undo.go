package actions

import (
	"fmt"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/Tools/myfetch"
	"github.com/Hana-ame/fedi-antenna/activitypub/actions"
	"github.com/Hana-ame/fedi-antenna/activitypub/model"
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/dao"
)

func UndoFollow(actor, id string) error {
	o := &activitypub.Follow{
		ID: id,
	}
	if err := dao.Read(&o); err != nil {
		return err
	}
	if o.Actor != actor {
		return fmt.Errorf("UndoFollow: %s != %s", o.Actor, actor)
	}

	if err := Undo(o); err != nil {
		return err
	}

	return nil
}

func UndoBlock(actor, id string) error {
	var o *activitypub.Block

	if err := dao.Read(&o); err != nil {
		return err
	}
	if o.Actor != actor {
		return fmt.Errorf("UndoBlock: %s != %s", o.Actor, actor)
	}

	if err := Undo(o); err != nil {
		return err
	}

	return nil
}

// all activitypub id url strings
func Undo(object activitypub.Sendable) error {

	o := &model.Undo{
		Object: object,
	}
	o.Autofill()
	dao.Create(o)

	body, err := myfetch.BuildJsonReader(o)
	if err != nil {
		return err
	}

	user, err := core.ReadActivitypubUserByID(o.GetObject())
	if err != nil {
		return err
	}
	resp, err := actions.Fetch(http.MethodPost, user.Inbox, nil, body)
	if err != nil {
		return err
	}

	_ = resp // todo?

	return nil
}
