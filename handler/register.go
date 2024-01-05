package handler

import (
	"os"

	tools "github.com/Hana-ame/fedi-antenna/Tools"
	"github.com/Hana-ame/fedi-antenna/antenna/model"
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
		Published:         utils.TimestampToRFC3339(utils.Now()),
		PrivateKeyPem:     utils.MarshalPrivateKey(pk),
		PublicKeyPem:      utils.MarshalPublicKey(&pk.PublicKey),
	}

	err := dao.Create(user)

	return err
}
