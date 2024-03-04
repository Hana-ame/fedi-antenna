package handler

import (
	"log"

	mastodon "github.com/Hana-ame/fedi-antenna/mastodon/dao"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

// id is a number in string
func Block_account(id, actor string) (*entities.Relationship, error) {
	tx := mastodon.DB.Begin()

	acct, err := mastodon.ReadAccount(tx, id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		relationship, _ := mastodon.ReadRelationship(tx, acct.Uri, actor)
		return relationship, err
	}

	if err := mastodon.Block(tx, acct.Uri, actor); err != nil {
		log.Println(err)
		tx.Rollback()
		relationship, _ := mastodon.ReadRelationship(tx, acct.Uri, actor)
		return relationship, err
	}

	relationship, err := mastodon.ReadRelationship(tx, acct.Uri, actor)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return relationship, err
	}

	tx.Commit()

	return relationship, tx.Error
}

func Unblock_account(id, actor string) (*entities.Relationship, error) {
	tx := mastodon.DB.Begin()

	acct, err := mastodon.ReadAccount(tx, id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		relationship, _ := mastodon.ReadRelationship(tx, acct.Uri, actor)
		return relationship, err
	}

	if err := mastodon.Unblock(tx, acct.Uri, actor); err != nil {
		log.Println(err)
		tx.Rollback()
		relationship, _ := mastodon.ReadRelationship(tx, acct.Uri, actor)
		return relationship, err
	}

	relationship, err := mastodon.ReadRelationship(tx, acct.Uri, actor)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return relationship, err
	}

	tx.Commit()

	return relationship, tx.Error
}
