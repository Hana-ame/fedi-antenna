package dao

import (
	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
	"gorm.io/gorm"
)

func DeleteReblog(tx *gorm.DB, status *entities.Status) error
