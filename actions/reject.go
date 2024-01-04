package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/activitypub/fetch"
	"github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

// all activitypub id url strings
func Reject(actor, object string) error {
	_, host := utils.ParseNameAndHost(actor)
	id := utils.GenerateObjectID("reject", host)

	followObj := &model.Follow{
		ID: object,
	}
	dao.Read(&followObj)
	followObj.Autofill()

	if actor != followObj.GetObject() {
		return fmt.Errorf("%s != %s", actor, followObj.GetObject())
	}

	o := &model.Reject{
		ID:     id,
		Object: followObj,
	}
	o.Autofill()

	dao.Create(o)

	body, err := json.Marshal(o)
	if err != nil {
		return err
	}

	user, err := core.ReadActivitypubUserByID(followObj.GetActor())
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
