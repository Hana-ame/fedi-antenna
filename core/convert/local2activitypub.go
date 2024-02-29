package convert

import (
	"log"

	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func ToActivityPubUser(activitypubID string) *activitypub.User {
	localuser := &model.LocalUser{
		ActivitypubID: activitypubID,
	}
	if err := dao.Read(dao.DB(), localuser); err != nil {
		log.Printf("%s", err.Error())
		return nil
	}
	acct := &entities.Account{
		Uri: activitypubID,
	}
	if err := dao.Read(dao.DB(), acct); err != nil {
		log.Printf("%s", err.Error())
		return nil
	}

	pk, err := utils.ParsePrivateKey(localuser.PrivateKeyPem)
	if err != nil {
		log.Printf("%s", err.Error())
		return nil
	}

	user := &activitypub.User{
		ID:                        activitypubID,
		Name:                      acct.DisplayName,
		Summary:                   acct.Note,
		ManuallyApprovesFollowers: acct.Locked,
		Discoverable:              utils.ParsePointerToBool(acct.Discoverable, false),
		Published:                 acct.CreatedAt,
		AlsoKnownAs:               localuser.AlsoKnownAs,
		PublicKey:                 &activitypub.PublicKey{ID: activitypubID + "#main-key", Owner: activitypubID, PublicKeyPem: utils.MarshalPublicKey(&pk.PublicKey)},
		Tag:                       nil,
		Attachment:                nil,
		Icon:                      ToActivityPubImage(acct.Avatar),
	}
	user.Autofill()
	return user
}

func ToActivityPubImage(imageURL string) *activitypub.Image {
	if imageURL == "" {
		return nil
	}
	image := &activitypub.Image{
		URL: imageURL,
	}
	dao.Read(dao.DB(), image)
	return image
}
