package handler

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Register(
	id string,
	username string,
	host string,
	ActivitypubID string,
) (err error) {
	tx := dao.DB().Begin()

	acct := &entities.Account{
		Id:       id,
		Username: username,
		Acct:     utils.ParseAcctStr(username, host),
		Url:      utils.NameAndHost2ProfileUrl(username, host),
		Uri:      ActivitypubID,
	}

	if err = dao.Create(tx, acct); err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	return nil
}
