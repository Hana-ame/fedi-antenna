package handler

import (
	"log"
	"os"
	"strconv"

	tools "github.com/Hana-ame/fedi-antenna/Tools"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	controller "github.com/Hana-ame/fedi-antenna/mastodon/controller/accounts/model"
	mastodon "github.com/Hana-ame/fedi-antenna/mastodon/dao"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Register_an_account(
	id string,
	host string,
	o *controller.Register_an_account,
) error {

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

	ctx := dao.Begin()
	if err := dao.Create(ctx, user); err != nil {
		log.Println(err)
		ctx.Rollback()
		return err
	}

	mtx := mastodon.DB.Begin()
	if err := mastodon.DB.Create(mtx, acct); err != nil {
		log.Println(err)
		mtx.Rollback()
		ctx.Rollback()
		return err
	}

	mtx.Commit()
	ctx.Commit()

	return nil
}

func Get_account(id, actor string) (*entities.Account, error) {
	tx := mastodon.DB.Begin()

	acct, err := mastodon.ReadAccount(tx, id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return acct, err
	}

	tx.Commit()

	return acct, tx.Error
}
