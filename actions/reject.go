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
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

// 'actor' will reject the request from 'object'
func Reject(object, actor string) error {
	lr := &model.LocalRelation{}
	if tx := dao.Where("Actor = ? AND Object = ?", object, actor).First(&lr); tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}

	if lr.Type != activitypub.TypeFollow { // should not.
		return fmt.Errorf("lr.Type != Follow")
	}
	_, host := utils.ParseNameAndHost(lr.Object)
	id := utils.GenerateObjectID("reject", host)

	follow := convert.ToActivityPubFollow(lr)
	follow.Autofill()

	o := &activitypub.Reject{
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

	_ = resp // todo?

	return nil
}
