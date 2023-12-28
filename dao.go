package core

import (
	"github.com/Hana-ame/fedi-antenna/activitypub/actions"
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
		user.Autofill()
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

func ReadPublicKeyByOwner(id string) (pk *activitypub.PublicKey, err error) {
	pk = &activitypub.PublicKey{
		Owner: id,
	}
	err = dao.Read(pk)
	if err == nil {
		return
	}
	var user *activitypub.User
	user, err = actions.FetchUserByID(id)
	if err != nil {
		return
	}
	dao.Create(user)

	pk = user.PublicKey

	return
}
