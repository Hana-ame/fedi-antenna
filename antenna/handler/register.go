package handler

import (
	"log"
	"os"
	"strconv"

	tools "github.com/Hana-ame/fedi-antenna/Tools"
	"github.com/Hana-ame/fedi-antenna/antenna/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/handler"
)

func Register(o *model.Register) error {
	id := utils.NameAndHost2ProfileUrlActivitypubID(o.Username, o.Host)
	pk := utils.GeneratePrivateKey()
	now := utils.NewTimestamp()

	user := &core.LocalUser{
		Email:         o.Email,
		PasswdHash:    tools.Hash(o.Passwd, os.Getenv("SALT")),
		ActivitypubID: id,
		AccountID:     strconv.Itoa(int(now)),
		Username:      o.Username,
		Host:          o.Host,
		CreatedAt:     now,
		PrivateKeyPem: utils.MarshalPrivateKey(pk),
		// PublicKeyPem:  utils.MarshalPublicKey(&pk.PublicKey),
	}

	if err := handler.Register(
		user.AccountID,
		user.Username,
		user.Host,
		id,
	); err != nil {
		return err
	}
	// acct := &entities.Account{
	// 	Id:       user.AccountID,
	// 	Username: user.Username,
	// 	Acct:     utils.ParseAcctStr(user.Username, user.Host),
	// 	Url:      utils.ParseProfileUrl(user.Username, user.Host),
	// 	Uri:      id,
	// }

	tx := dao.DB().Begin()
	if err := dao.Create(tx, user); err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}
	// if err := dao.Create(tx, acct); err != nil {
	// 	log.Println(err)
	// 	tx.Rollback()
	// 	return err
	// }

	return nil
}
