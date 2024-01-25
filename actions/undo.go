package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/actions/fetch"
	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	model "github.com/Hana-ame/fedi-antenna/core/model"
)

func UndoFollow(object, actor string) error {
	lr := &model.LocalRelation{}
	if tx := dao.Where("Object = ? AND Actor = ?", object, actor).First(&lr); tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	if lr.Type != model.RelationTypeFollow {
		return fmt.Errorf("not follow")
	}
	if lr.Status != model.RelationStatusPadding {
		return fmt.Errorf("not padding")
	}

	o := &activitypub.Follow{
		ID:     lr.ID,
		Actor:  actor,
		Object: object,
	}
	// if tx := dao.Where("Object = ? AND Actor = ?", object, actor).First(&o); tx.Error != nil {
	// 	log.Println(tx.Error)
	// 	return tx.Error
	// }

	o.Autofill()
	o.ClearContext()
	if err := Undo(o); err != nil {
		return err
	}

	dao.Delete(lr)

	return nil
}

func UndoBlock(object, actor string) error {
	lr := &model.LocalRelation{}
	if tx := dao.Where("Object = ? AND Actor = ?", object, actor).First(&lr); tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}

	o := &activitypub.Block{
		ID:     lr.ID,
		Actor:  actor,
		Object: object,
	}
	if lr.Type != model.RelationTypeBlock {
		return fmt.Errorf("not block")
	}

	o.Autofill()
	o.ClearContext()
	if err := Undo(o); err != nil {
		return err
	}

	dao.Delete(lr)

	return nil
}

// all activitypub id url strings
func Undo(object activitypub.Object) error {

	o := &activitypub.Undo{
		Object: object,
	}
	o.Autofill()
	dao.Create(o)
	// dao.Delete(object)

	body, err := json.Marshal(o)
	if err != nil {
		return err
	}

	resp, err := fetch.FetchWithSign(
		o.GetActor(),
		http.MethodPost, o.GetEndpoint(), nil, body,
	)
	if err != nil {
		return err
	}

	_ = resp // todo?
	// log.Printf("%s", body)
	return nil
}
