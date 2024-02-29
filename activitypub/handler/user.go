package handler

import (
	"github.com/Hana-ame/fedi-antenna/core/dao"
)

func DeletePerson(id string) error {
	return dao.DeletePerson(id)
}
