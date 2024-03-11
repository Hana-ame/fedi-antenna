package handler

import (
	"github.com/Hana-ame/fedi-antenna/core/utils"
	"github.com/Hana-ame/fedi-antenna/db"
	"github.com/Hana-ame/fedi-antenna/mastodon/dao"
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

func View_public_timeline(
	Authorization,
	local,
	remote,
	only_media,
	max_id,
	since_id,
	min_id,
	limit string,
) ([]*entities.Status, error) {

	tx := db.Begin()

	statuses, err := dao.ReadPublicStatus(tx, utils.Atoi(limit, 20))

	if err != nil {
		logE(err)
		tx.Rollback()
		return statuses, err
	}

	return statuses, tx.Error
}
