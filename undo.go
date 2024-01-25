package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/actions/fetch"
	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
)

func UndoFollow(actor, id string) error {
	o := &activitypub.Follow{
		ID: id,
	}
	if err := dao.Read(o); err != nil {
		return err
	}
	if o.Actor != actor {
		return fmt.Errorf("UndoFollow: %s != %s", o.Actor, actor)
	}
	o.Autofill()
	o.ClearContext()
	if err := Undo(o); err != nil {
		return err
	}

	return nil
}

func UndoBlock(actor, id string) error {
	var o *activitypub.Block

	if err := dao.Read(o); err != nil {
		return err
	}
	if o.Actor != actor {
		return fmt.Errorf("UndoBlock: %s != %s", o.Actor, actor)
	}
	o.Autofill()
	o.ClearContext()
	if err := Undo(o); err != nil {
		return err
	}

	return nil
}

// all activitypub id url strings
func Undo(object activitypub.Object) error {

	o := &activitypub.Undo{
		Object: object,
	}
	o.Autofill()
	dao.Create(o)
	// dao.Delete(object)

	body, err := json.Marshal(o)
	if err != nil {
		return err
	}

	resp, err := fetch.FetchWithSign(
		o.GetActor(),
		http.MethodPost, o.GetEndpoint(), nil, body,
	)
	if err != nil {
		return err
	}

	_ = resp // todo?
	// log.Printf("%s", body)
	return nil
}
