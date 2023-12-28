package actions

import (
	"encoding/json"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/activitypub/actions"
	"github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

// all activitypub id url strings
func Block(actor, object string) error {
	_, host := utils.ParseNameAndHost(actor)
	id := utils.GenerateObjectID("block", host)
	o := &model.Block{
		ID:     id,
		Actor:  actor,
		Object: object,
	}
	o.Autofill()
	err := dao.Create(o)
	if err != nil {
		return err
	}
	body, err := json.Marshal(o)
	if err != nil {
		return err
	}
	user, err := core.ReadActivitypubUserByID(object)
	if err != nil {
		return err
	}
	local, err := core.ReadActivitypubUserByID(actor)
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

	return nil
}
