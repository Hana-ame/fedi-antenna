package dao

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"gorm.io/gorm"
)

func CreateFavourite(tx *gorm.DB, object, actor string) error {
	status := &entities.Status{Uri: object}
	if err := DB.Read(tx, status); err != nil {
		log.Println(err)
		return err
	}

	favourite := &Favourite{
		Actor:  actor,
		Object: object,
		Owner:  status.AttributedTo,
	}
	if err := DB.Create(tx, favourite); err != nil {
		log.Println(err)
		return err
	}

	tx.Commit()

	return tx.Error
}

func DeleteFavourite(tx *gorm.DB, object, actor string) error {
	favourite := &Favourite{
		Actor:     actor,
		Object:    object,
		DeletedAt: utils.NewTimestamp(true),
	}
	if err := DB.Update(tx, favourite); err != nil {
		log.Println(err)
		return err
	}

	tx.Commit()

	return tx.Error
}
