package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/actions/fetch"
	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
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
			log.Println(err)
			return err
		}
	}

	if lr.Type != activitypub.TypeFollow { // should not.
		return fmt.Errorf("lr.Type != Follow")
	}
	_, host := utils.ActivitypubID2NameAndHost(lr.Object)
	id := utils.NewObjectID("accept", host)

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
	lr.Status = model.RelationStatusAccepted
	dao.Update(lr)

	return nil
}
