package actions

import (
	"encoding/json"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/activitypub/fetch"
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

	resp, err := fetch.FetchWithSign(
		actor,
		http.MethodPost, user.Inbox, nil, body,
	)
	if err != nil {
		return err
	}

	_ = resp // todo?

	return nil
}
