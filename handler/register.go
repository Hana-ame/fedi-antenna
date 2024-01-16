package handler

import (
	"os"

	tools "github.com/Hana-ame/fedi-antenna/Tools"
	"github.com/Hana-ame/fedi-antenna/antenna/model"
	"github.com/Hana-ame/fedi-antenna/core/convert"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

func Register(o *model.Register) error {
	id := utils.ParseActivitypubID(o.Username, o.Host)
	pk := utils.GeneratePrivateKey()
	user := &core.LocalUser{
		Email:             o.Email,
		PasswdHash:        tools.Hash(o.Passwd, os.Getenv("SALT")),
		ID:                id,
		PreferredUsername: o.Username,
		Published:         utils.Now(),
		PrivateKeyPem:     utils.MarshalPrivateKey(pk),
		PublicKeyPem:      utils.MarshalPublicKey(&pk.PublicKey),
	}

	if err := dao.Create(user); err != nil {
		return err
	}

	activitypubUser := convert.ToActivityPubUser(user)
	if err := dao.Create(activitypubUser); err != nil {
		return err
	}

	return nil
}
