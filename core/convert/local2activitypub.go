package convert

import (
	"log"

	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func AccountToActivitypubUser(acct *entities.Account, localuser *model.LocalUser) *activitypub.User {

	pk, err := utils.ParsePrivateKey(localuser.PrivateKeyPem)
	if err != nil {
		log.Printf("%s", err.Error())
		return nil
	}

	icon := &activitypub.Image{
		URL: acct.Avatar,
	}
	if err := dao.Read(dao.DB(), icon); err != nil {
		log.Printf("%s", err.Error())
		return nil
	}

	user := &activitypub.User{
		ID:                        localuser.ActivitypubID,
		Name:                      acct.DisplayName,
		Summary:                   acct.Note,
		ManuallyApprovesFollowers: acct.Locked,
		Discoverable:              utils.ParsePointerToBool(acct.Discoverable, false),
		Published:                 acct.CreatedAt,
		AlsoKnownAs:               localuser.AlsoKnownAs,
		PublicKey: &activitypub.PublicKey{
			ID:           localuser.ActivitypubID + "#main-key",
			Owner:        localuser.ActivitypubID,
			PublicKeyPem: utils.MarshalPublicKey(&pk.PublicKey)},
		Tag:        nil,
		Attachment: nil,
		Icon:       icon,
	}
	user.Autofill()
	return user
}

func LocalRelationToFollow(o *core.LocalRelation) *activitypub.Follow {
	follow := &activitypub.Follow{
		ID:     o.ID,
		Actor:  o.Actor,
		Object: o.Object,
	}
	follow.Autofill()
	return follow
}
