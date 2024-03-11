package handler

import (
	"os"
	"strconv"

	tools "github.com/Hana-ame/fedi-antenna/Tools"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/db"
	"github.com/Hana-ame/fedi-antenna/mastodon/controller/accounts/model"
	"github.com/Hana-ame/fedi-antenna/mastodon/dao"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Register_an_account(
	id string,
	host string,
	o *model.Register_an_account,
) error {

	activitypubID := utils.NameAndHost2ActivitypubID(o.Username, host)
	pk := utils.GeneratePrivateKey()
	now := utils.Timestamp(false)

	user := &dao.MyAccount{
		Email:         o.Email,
		PasswdHash:    tools.Hash(o.Password, os.Getenv("SALT")),
		ActivitypubID: activitypubID,
		AccountID:     strconv.Itoa(int(now)),
		Username:      o.Username,
		Host:          host,
		CreatedAt:     now,
		PrivateKeyPem: utils.MarshalPrivateKey(pk),
	}

	acct := &entities.Account{
		Id:       strconv.Itoa(int(now)),
		Username: o.Username,
		Acct:     utils.UsernameAndHost2Account(o.Username, host),
		Url:      utils.NameAndHost2ProfileUrl(o.Username, host),
		Uri:      activitypubID,
	}

	tx := db.Begin()

	if err := dao.Create(tx, acct); err != nil {
		logE(err)
		tx.Rollback()
		return err
	}

	if err := dao.Create(tx, user); err != nil {
		logE(err)
		tx.Rollback()
		return err
	}

	tx.Commit()

	return tx.Error
}

func Get_account(id, actor string) (*entities.Account, error) {

	tx := db.Begin()

	acct := &entities.Account{Id: id}
	_, err := dao.ReadAccount(tx, acct)
	if err != nil {
		logE(err)
		tx.Rollback()
		return acct, err
	}

	tx.Commit()

	return acct, tx.Error
}
