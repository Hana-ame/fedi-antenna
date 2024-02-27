package core

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func CreateStatus(status *entities.Status) error {
	if err := dao.CreateStatus(status); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func DeleteStatus(status *entities.Status) error {
	if err := dao.DeleteStatus(status); err != nil {
		// if there is no note
		// it is actually not found error.
		log.Println(err)
		return err
	}

	return nil
}
