package actions

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/activitypub/fetch"
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	model "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/core/utils/convert"
)

// all activitypub id url strings
func Accept(actor, object string) error {
	_, host := utils.ParseNameAndHost(actor)
	id := utils.GenerateObjectID("accept", host)

	lr := &model.LocalRelation{
		Actor:  object,
		Object: actor,
	}
	if err := dao.Read(&lr); err != nil {
		return err
	}
	if lr.Type != activitypub.TypeFollow { // should not.
		return fmt.Errorf("lr.Type != activitypub.TypeFollow")
	}
	follow := convert.ToActivityPubFollow(lr)
	follow.Autofill()

	if actor != follow.Object { // never.
		return fmt.Errorf("%s != %s", actor, follow.Object)
	}

	o := &activitypub.Accept{
		ID:     id,
		Actor:  actor,
		Object: follow,
	}
	o.Autofill()

	dao.Create(o)

	body, err := json.Marshal(o)
	if err != nil {
		return err
	}

	user, err := core.ReadActivitypubUserByID(follow.GetActor())
	if err != nil {
		return err
	}

	resp, err := fetch.FetchWithSign(
		follow.Object,
		http.MethodPost, user.Inbox, nil, body,
	)
	if err != nil {
		return err
	}

	_ = resp

	return nil
}
