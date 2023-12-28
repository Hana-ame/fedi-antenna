package core

import (
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

func ReadActivitypubUser(name, host string) (user *activitypub.User, err error) {
	host = Host(host)

	user = &activitypub.User{
		ID: utils.ParseActivitypubID(name, host),
	}

	err = dao.Read(user)
	if err == nil {
		user.Patch()
		user.PublicKey = &activitypub.PublicKey{
			Owner: user.ID,
		}
		dao.Read(user.PublicKey)
		user.Icon = &activitypub.Image{
			URL: user.IconURL,
		}
		dao.Read(user.Icon)
	}
	return
}
