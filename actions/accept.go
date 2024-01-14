package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/activitypub/fetch"
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/core"
	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	model "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

// all activitypub id url strings
func Accept(lr *model.LocalRelation, shouldRead bool) error {
	if lr == nil {
		log.Printf("nothing passed\n")
		return fmt.Errorf("nothing passed")
	}
	if shouldRead {
		if err := dao.Read(&lr); err != nil {
			return err
		}
	}

	if lr.Type != activitypub.TypeFollow { // should not.
		return fmt.Errorf("lr.Type != Follow")
	}
	_, host := utils.ParseNameAndHost(lr.Object)
	id := utils.GenerateObjectID("accept", host)

	follow := convert.ToActivityPubFollow(lr)
	follow.Autofill()

	o := &activitypub.Accept{
		ID:     id,
		Actor:  follow.Object,
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
