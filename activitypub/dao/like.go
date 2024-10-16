package dao

import (
	"github.com/Hana-ame/fedi-antenna/activitypub/dao/model"
)

func Like(id, actor, object string, createdAt int64) (err error) {

	like := &model.Like{
		ID:     id,
		Actor:  actor,
		Object: object,

		CratedAt: createdAt,
	}

	err = Create(like)

	return
}

func UndoLike(id, actor, object string, deletedAt int64) (err error) {

	like := &model.Like{
		ID:     id,
		Actor:  actor,
		Object: object,

		DeletedAt: deletedAt,
	}

	err = Update(like)

	return
}
