package core

import (
	"github.com/Hana-ame/fedi-antenna/activitypub/actions"
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	webfinger "github.com/Hana-ame/fedi-antenna/webfinger/actions"
)

// convert an altername to it's origin.
func Host(alias string) string {
	host, exist := utils.AliasMap[alias]
	if exist {
		return host
	}
	return alias
}

// find in local
// if not found then fetch from remote
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

func ReadActivitypubUser(name, host string) (user *activitypub.User, err error) {
	host = Host(host)
	id, err := webfinger.GetUserIdFromAcct(utils.ParseAcctStr(name, host))
	if err != nil {
		return
	}
	user, err = ReadActivitypubUserByID(id)
	return
}

func ReadActivitypubUserByID(id string) (user *activitypub.User, err error) {
	if user, err = dao.ReadActivitypubUser(id); err == nil {
		// this will lead to hell...
		// local user and remote user should be in different tables.
		user.Type = "Person"
		return
	}
	if user, err = actions.FetchUserByID(id); err != nil {
		return
	}
	dao.Create(user)
	return
}
