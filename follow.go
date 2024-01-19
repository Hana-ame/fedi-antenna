package actions

import (
	"encoding/json"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/actions/fetch"
	"github.com/Hana-ame/fedi-antenna/actions/model"
	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

// only for local user follow another user
// actor is the activityID of yourself.
// object is the activityID of object.
// all activitypub id url strings
func Follow(actor, object string) error {
	_, host := utils.ParseNameAndHost(actor)
	id := utils.GenerateObjectID("follow", host)
	o := &model.Follow{
		ID:     id,
		Actor:  actor,
		Object: object,
	}
	o.Autofill()
	dao.Create(o)
	body, err := json.Marshal(o)
	if err != nil {
		return err
	}

	user, err := core.ReadActivitypubUserByID(o.Object)
	if err != nil {
		return err
	}

	resp, err := fetch.FetchWithSign(
		o.GetActor(),
		http.MethodPost, user.Inbox, nil, body,
	)
	if err != nil {
		return err
	}
	_ = user
	_ = resp // todo?
	// log.Printf("%s", body)

	return nil
}
