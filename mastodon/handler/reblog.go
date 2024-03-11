package handler

import (
	"log"

	controller "github.com/Hana-ame/fedi-antenna/mastodon/controller/statuses/model"
	mastodon "github.com/Hana-ame/fedi-antenna/mastodon/dao"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Boost_a_status(id string, actor string, o *controller.Boost_a_status) (*entities.Status, error) {
	tx := mastodon.DB.Begin()

	status := &entities.Status{
		Id: id,
	}
	if err := mastodon.DB.Read(tx, status); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return nil, err
	}

	if err := mastodon.CreateReblog(tx, status.Uri, actor, o.Visibility); err != nil {
		log.Println(err)
		tx.Rollback()
		return status, err
	}

	tx.Commit()

	return status, tx.Error
}

func Undo_boost_of_a_status(id string, actor string) (*entities.Status, error) {
	tx := mastodon.DB.Begin()

	status := &entities.Status{
		Id: id,
	}
	if err := mastodon.DB.Read(tx, status); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return nil, err
	}

	if err := mastodon.DeleteReblog(tx, status.Uri, actor); err != nil {
		log.Println(err)
		tx.Rollback()
		return status, err
	}

	tx.Commit()

	return status, tx.Error
}
