package handler

import (
	"log"

	mastodon "github.com/Hana-ame/fedi-antenna/mastodon/dao"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

// id is the mastodon published which is timestamp in us
func Favourite_a_status(id string, actor string) (*entities.Status, error) {
	tx := mastodon.DB.Begin()

	status := &entities.Status{
		Id: id,
	}
	if err := mastodon.DB.Read(tx, status); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return nil, err
	}

	if err := mastodon.CreateFavourite(tx, status.Uri, actor); err != nil {
		log.Println(err)
		tx.Rollback()
		return status, err
	}

	tx.Commit()

	return status, tx.Error
}

func Undo_favourite_of_a_status(id string, actor string) (*entities.Status, error) {
	tx := mastodon.DB.Begin()

	status := &entities.Status{
		Id: id,
	}
	if err := mastodon.DB.Read(tx, status); err != nil {
		log.Printf("%s", err.Error())
		tx.Rollback()
		return nil, err
	}

	if err := mastodon.DeleteFavourite(tx, status.Uri, actor); err != nil {
		log.Println(err)
		tx.Rollback()
		return status, err
	}

	tx.Commit()

	return status, tx.Error
}
