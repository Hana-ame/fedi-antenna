package dao

import (
	"github.com/Hana-ame/fedi-antenna/activitypub/dao/model"
	"github.com/Hana-ame/fedi-antenna/utils"
)

func Follow(id, actor, object string, createdAt int64) (err error) {

	follow := &model.Follow{
		ID:     id,
		Actor:  actor,
		Object: object,

		CreatedAt: createdAt,
	}

	err = Create(follow)

	return
}

func UndoFollow(id, actor, object string, deletedAt int64) (err error) {

	follow := &model.Follow{
		ID:     id,
		Actor:  actor,
		Object: object,

		DeletedAt: deletedAt,
	}

	err = Update(follow)

	return
}

func Accept(id, actor, object string, acceptedAt int64) (err error) {

	follow := &model.Follow{
		ID:     id,
		Actor:  actor,
		Object: object,

		AcceptedAt: acceptedAt,
	}

	err = Update(follow)

	return
}

func Reject(id, actor, object string, rejectedAt int64) (err error) {

	follow := &model.Follow{
		ID:     id,
		Actor:  actor,
		Object: object,

		RejectedAt: rejectedAt,
	}

	err = Update(follow)

	return
}

func Block(id, actor, object string, createdAt int64) (err error) {

	block := &model.Block{
		ID:     id,
		Actor:  actor,
		Object: object,

		CreatedAt: createdAt,
	}

	err = Create(block)
	// if success, all requests are deleted.
	if err != nil {
		// TBD: not checked.
		db.Exec(`UPDATE flollows SET blocked_at = ? WHERE actor = ? AND object = ?`, utils.Now(), actor, object)
		db.Exec(`UPDATE flollows SET blocked_at = ? WHERE actor = ? AND object = ?`, utils.Now(), object, actor)
	}

	return
}

func UndoBlock(id, actor, object string, deleteAt int64) (err error) {
	block := &model.Block{
		ID:     id,
		Actor:  actor,
		Object: object,

		DeletedAt: deleteAt,
	}

	err = Create(block)

	return
}
