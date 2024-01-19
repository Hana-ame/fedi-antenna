package actions

import (
	"encoding/json"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/actions/fetch"
	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

// 未完成
func Like(actor, object string) error {
	_, host := utils.ParseNameAndHost(actor)
	id := utils.GenerateObjectID("like", host)

	o := &activitypub.Like{
		ID:     id,
		Actor:  actor,
		Object: object,
	}
	o.Autofill()

	ln := &core.LocalNote{
		ID: object,
	}
	if err := dao.Read(ln); err != nil {
		return err
	}
	// user := core.ReadActivitypubUserByAccount(ln.AttributedTo)
	user := &core.LocalNote{}

	body, err := json.Marshal(o)
	if err != nil {
		return err
	}

	resp, err := fetch.FetchWithSign(
		o.GetActor(),
		http.MethodPost, "user.Inbox", nil, body,
	)
	if err != nil {
		return err
	}
	_ = user
	_ = resp // todo?
	// log.Printf("%s", body)
	return err
}
