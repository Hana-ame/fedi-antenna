package handler

import (
	"log"
	"os"
	"strconv"

	tools "github.com/Hana-ame/fedi-antenna/Tools"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	mastodon "github.com/Hana-ame/fedi-antenna/mastodon/controller/accounts/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Register_an_account(
	id string,
	host string,
	o *mastodon.Register_an_account,
) error {
	tx := dao.DB().Begin()

	activitypubID := utils.NameAndHost2ActivitypubID(o.Username, host)
	pk := utils.GeneratePrivateKey()
	now := utils.NewTimestamp(false)

	user := &model.LocalUser{
		Email:         o.Email,
		PasswdHash:    tools.Hash(o.Password, os.Getenv("SALT")),
		ActivitypubID: activitypubID,
		AccountID:     strconv.Itoa(int(now)),
		Username:      o.Username,
		Host:          host,
		CreatedAt:     now,
		PrivateKeyPem: utils.MarshalPrivateKey(pk),
		// PublicKeyPem:  utils.MarshalPublicKey(&pk.PublicKey),
	}

	acct := &entities.Account{
		Id:       strconv.Itoa(int(now)),
		Username: o.Username,
		Acct:     utils.UsernameAndHost2Account(o.Username, host),
		Url:      utils.NameAndHost2ProfileUrl(o.Username, host),
		Uri:      activitypubID,
	}

	if err := dao.Create(tx, user); err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}
	if err := dao.Create(tx, acct); err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}

func Get_account(id, actor string) (*entities.Account, error) {
	tx := dao.Begin()

	acct := &entities.Account{
		Id: id,
	}
	if err := dao.Read(tx, acct); err != nil {
		log.Println(err)
		tx.Rollback()
		return acct, err
	}

	tx.Commit()

	return acct, tx.Error
}
