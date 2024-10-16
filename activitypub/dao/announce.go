package dao

import (
	"github.com/Hana-ame/fedi-antenna/activitypub/dao/model"
)

func Announce(id, actor, object, visibility string, createdAt int64) (err error) {

	like := &model.Announce{
		ID:         id,
		Actor:      actor,
		Object:     object,
		Visibility: visibility,

		CreatedAt: createdAt,
	}

	err = Create(like)

	return
}

func UndoAnnounce(id, actor, object string, deletedAt int64) (err error) {

	like := &model.Announce{
		ID:     id,
		Actor:  actor,
		Object: object,

		DeletedAt: deletedAt,
	}

	err = Update(like)

	return
}
