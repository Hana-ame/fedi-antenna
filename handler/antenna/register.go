package handler

import (
	"os"

	tools "github.com/Hana-ame/fedi-antenna/Tools"
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/antenna/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

func Register(o *model.Register) error {
	perferedUsername := o.Username
	host := o.Host
	email := o.Email
	passwd := o.Passwd

	id := utils.ParseActivitypubID(perferedUsername, host)
	pk := activitypub.NewPublicKey(id)
	published := utils.TimestampToRFC3339(utils.Now())
	user := &core.User{
		PreferredUsername: perferedUsername,
		ID:                id,
		Email:             email,
		PasswdHash:        tools.Hash(passwd, os.Getenv("SALT")),
		Published:         published,
		PublicKey:         pk,
	}

	err := dao.Create(user)

	return err
}
