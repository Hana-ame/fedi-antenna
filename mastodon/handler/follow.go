package handler

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
	accounts "github.com/Hana-ame/fedi-antenna/mastodon/controller/accounts/model"
	mastodon "github.com/Hana-ame/fedi-antenna/mastodon/dao"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func Follow_account(id, actor string, o *accounts.Follow_account) (*entities.Relationship, error) {
	tx := mastodon.DB.Begin()

	acct, err := mastodon.ReadAccount(tx, id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		relationship, _ := mastodon.ReadRelationship(tx, acct.Uri, actor)
		return relationship, err
	}

	if err := mastodon.Follow(tx, acct.Uri, actor); err != nil {
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

func Unfollow_account(id, actor string) (*entities.Relationship, error) {
	tx := mastodon.DB.Begin()

	acct, err := mastodon.ReadAccount(tx, id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		relationship, _ := mastodon.ReadRelationship(tx, acct.Uri, actor)
		return relationship, err
	}

	if err := mastodon.Unfollow(tx, acct.Uri, actor); err != nil {
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

// TODO
func View_pending_follow_requests(
	Authorization,
	max_id,
	since_id,
	limit string,
) (accts []*entities.Account, err error) {
	// todo
	// utils.Atoi(max_id, 0)
	tx := dao.DB().Limit(utils.Atoi(limit, 25)).Find(&accts)
	if tx.Error != nil {
		return accts, tx.Error
	}
	return accts, nil
}

func Accept_follow_request(id, actor string) (*entities.Relationship, error) {
	tx := mastodon.DB.Begin()

	acct, err := mastodon.ReadAccount(tx, id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		relationship, _ := mastodon.ReadRelationship(tx, acct.Uri, actor)
		return relationship, err
	}

	if err := mastodon.Accept(tx, acct.Uri, actor); err != nil {
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

func Reject_follow_request(id, actor string) (*entities.Relationship, error) {
	tx := mastodon.DB.Begin()

	acct, err := mastodon.ReadAccount(tx, id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		relationship, _ := mastodon.ReadRelationship(tx, acct.Uri, actor)
		return relationship, err
	}

	if err := mastodon.Reject(tx, acct.Uri, actor); err != nil {
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
