package core

import (
	"log"

	actions "github.com/Hana-ame/fedi-antenna/actions/fetch"
	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	webfinger "github.com/Hana-ame/fedi-antenna/webfinger/actions"
)

// find in local
// if not found then fetch from remote
func ReadPublicKeyByOwner(id string) (pk *activitypub.PublicKey, err error) {
	pk = &activitypub.PublicKey{
		Owner: id,
	}
	err = dao.Where(*pk).Find(pk).Error
	if err == nil {
		return
	}
	var user *activitypub.User
	user, err = actions.FetchUserByID(id)
	if err != nil {
		log.Println(err)
		return
	}
	dao.Create(user)

	pk = user.PublicKey

	return
}

func ReadActivitypubUser(name, host string) (user *activitypub.User, err error) {
	host = Host(host)
	id, err := webfinger.GetUserIdFromAcct(utils.ParseAcctStr(name, host), false)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return
	}
	dao.Create(user)
	return
}
