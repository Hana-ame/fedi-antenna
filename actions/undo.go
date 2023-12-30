package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/activitypub/actions"
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
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
	o.Autofill()
	o.ClearContext()
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
	o.Autofill()
	o.ClearContext()
	if err := Undo(o); err != nil {
		return err
	}

	return nil
}

// all activitypub id url strings
func Undo(object activitypub.Sendable) error {

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

	user, err := core.ReadActivitypubUserByID(o.GetObject(), false)
	if err != nil {
		return err
	}
	local, err := core.ReadActivitypubUserByID(o.GetActor(), true)
	if err != nil {
		return err
	}
	pk, err := utils.ParsePrivateKey(local.PublicKey.PrivateKeyPem)
	if err != nil {
		return err
	}
	resp, err := actions.FetchWithSign(
		pk, local.PublicKey.ID,
		http.MethodPost, user.Inbox, nil, body,
	)
	if err != nil {
		return err
	}

	_ = resp // todo?
	fmt.Printf("%s", body)
	return nil
}
