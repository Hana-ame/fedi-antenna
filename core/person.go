package core

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/actions/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
)

func DeletePerson(id string) error {
	person := &model.User{
		ID: id,
	}
	if err := dao.Delete(person); err != nil {
		log.Println(err)
		return err
	}
	// delete all notify
	// delete all notes
	return nil
}

func CachePerson(person *model.User) error {
	if err := dao.Create(person); err != nil {
		if err := dao.Update(person); err != nil {
			return err
		}
		return err
	}
	return nil
}
