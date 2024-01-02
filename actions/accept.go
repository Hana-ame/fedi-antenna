package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/activitypub/actions"
	"github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

// all activitypub id url strings
func Accept(actor, object string) error {
	followObj := &model.Follow{
		ID: object,
	}
	dao.Read(&followObj)
	followObj.Autofill()

	if actor != followObj.GetObject() {
		return fmt.Errorf("%s != %s", actor, followObj.GetObject())
	}

	_, host := utils.ParseNameAndHost(actor)
	id := utils.GenerateObjectID("accept", host)
	o := &model.Accept{
		ID:     id,
		Object: followObj,
	}
	o.Autofill()

	// dao.Create(o)

	body, err := json.Marshal(o)
	if err != nil {
		return err
	}

	user, err := core.ReadActivitypubUserByID(followObj.GetActor())
	if err != nil {
		return err
	}

	_, err = actions.FetchWithSign(
		followObj.GetObject(),
		http.MethodPost, user.Inbox, nil, body,
	)
	if err != nil {
		return err
	}

	// _ = resp // todo?
	// _ = user
	// fmt.Printf("%s", body)
	return nil
}
