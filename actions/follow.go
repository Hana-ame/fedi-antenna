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

	user, err := core.ReadActivitypubUserByID(o.GetObject())
	if err != nil {
		return err
	}
	local, err := core.ReadActivitypubUserByID(o.GetActor())
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
	_ = user
	_ = resp // todo?
	fmt.Printf("%s", body)
	
	return nil
}
