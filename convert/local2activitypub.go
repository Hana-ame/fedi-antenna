package convert

import (
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

func ToActivityPubUser(o *core.LocalUser) *activitypub.User {
	// name, host := utils.ParseNameAndHost(o.ID)
	icon := &activitypub.Image{
		URL: o.IconURL,
	}
	if o.IconURL == "" {
		icon = nil
	} else {
		dao.Read(icon)
	}

	user := &activitypub.User{
		ID:                        o.ID,
		Name:                      o.Name,
		Summary:                   o.Summary,
		ManuallyApprovesFollowers: o.ManuallyApprovesFollowers,
		Discoverable:              o.Discoverable,
		Published:                 utils.TimestampToRFC3339(o.Published),
		AlsoKnownAs:               o.AlsoKnownAs,
		PublicKey:                 &activitypub.PublicKey{ID: o.ID + "#main-key", Owner: o.ID, PublicKeyPem: o.PublicKeyPem},
		Tag:                       nil,
		Attachment:                nil,
		Icon:                      icon,
	}
	user.Autofill()
	return user
}
