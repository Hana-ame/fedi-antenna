package inbox

import (
	"github.com/Hana-ame/fedi-antenna/activitypub/dao"
	"github.com/Hana-ame/fedi-antenna/activitypub/dao/model"
	"github.com/Hana-ame/fedi-antenna/utils"
)

func DeletePerson(id string) (err error) {
	user := &model.User{
		ID:        id,
		DeletedAt: utils.Now(),
	}

	err = dao.Update(user)

	return
}
