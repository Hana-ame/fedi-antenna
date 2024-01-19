package handler

import (
	"log"

	"github.com/Hana-ame/fedi-antenna/core/dao"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"

	"github.com/Hana-ame/fedi-antenna/mastodon/entities"
)

// id is the mastodon published which is timestamp in us
func Favourite_a_status(id string, actor string) (*entities.Status, error) {
	// published, err := strconv.Atoi(id)
	// if err != nil {
	// 	log.Printf("%s", err.Error())
	// 	return nil, err
	// }
	status := &entities.Status{
		Id: id,
	}
	if err := dao.Read(status); err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}
	_, host := utils.ParseNameAndHost(actor)

	// 考虑一下
	favouriteID := utils.GenerateObjectID("favourite", host)
	favourite := &core.LocalNotify{
		ID:     favouriteID,
		Actor:  actor,
		Object: status.Uri,
		Type:   core.NotifyTypeLike,
	}

	// acitiviypub
	// todo
	//

	if err := dao.Create(favourite); err != nil {
		log.Printf("%s", err.Error())
		return status, err
	}

	// mastodon
	return status, nil
}

func Undo_favourite_of_a_status(id string, actor string) (*entities.Status, error) {
	// published, err := strconv.Atoi(id)
	// if err != nil {
	// 	log.Printf("%s", err.Error())
	// 	return nil, err
	// }
	status := &entities.Status{
		Id: id,
	}
	if err := dao.Read(status); err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}
	favourite := &core.LocalNotify{
		Actor:  actor,
		Object: status.Uri,
		Type:   core.NotifyTypeLike,
	}
	if err := dao.Read(favourite); err != nil {
		log.Printf("%s", err.Error())
		return nil, err
	}

	// acitiviypub
	// todo
	//

	if err := dao.Delete(favourite); err != nil {
		log.Printf("%s", err.Error())
		return status, err
	}

	// mastodon
	return status, nil
}
