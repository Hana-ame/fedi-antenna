package handler

import "github.com/Hana-ame/fedi-antenna/core"

func DeletePerson(id string) error {
	return core.DeletePerson(id)
}
