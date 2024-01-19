package handler

import (
	"os"

	tools "github.com/Hana-ame/fedi-antenna/Tools"
	"github.com/Hana-ame/fedi-antenna/antenna/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Register(o *model.Register) error {
	id := utils.ParseActivitypubID(o.Username, o.Host)
	pk := utils.GeneratePrivateKey()
	now := utils.Now()
	user := &core.LocalUser{
		Email:         o.Email,
		PasswdHash:    tools.Hash(o.Passwd, os.Getenv("SALT")),
		ActivitypubID: id,
		AccountID:     utils.TimestampToRFC3339(now),
		Username:      o.Username,
		Host:          o.Host,
		CreatedAt:     now,
		PrivateKeyPem: utils.MarshalPrivateKey(pk),
		// PublicKeyPem:  utils.MarshalPublicKey(&pk.PublicKey),
	}

	acct := &entities.Account{
		Id:       utils.TimestampToRFC3339(user.CreatedAt),
		Username: user.Username,
		Acct:     utils.ParseAcctStr(user.Username, user.Host),
		Url:      utils.ParseProfileUrl(user.Username, user.Host),
		Uri:      user.AccountID,
	}
	if err := dao.Create(user); err != nil {
		return err
	}
	if err := dao.Create(acct); err != nil {
		return err
	}

	// activitypubUser := convert.ToActivityPubUser(user)
	// if err := dao.Create(activitypubUser); err != nil {
	// 	return err
	// }

	return nil
}
