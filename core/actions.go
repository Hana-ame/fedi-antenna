package core

import (
	"errors"

	"github.com/Hana-ame/fedi-antenna/activitypub"
	"github.com/Hana-ame/fedi-antenna/utils"
	"github.com/iancoleman/orderedmap"
)

func Follow(user, host, object string) (err error) {
	uuid := utils.NewUUIDString()
	actor := activitypub.APUserID(user, host)
	id := utils.ObjectId(uuid, host)
	followObj := activitypub.Follow(id, actor, object)

	// endpoint
	inbox := object + "/inbox"
	// sign object
	req, err := signObjectByUser(followObj, user, host, inbox)
	if err != nil {
		return err
	}
	_, err = fetchObjectWithRequest(req)

	return
}

func UndoFollow(user, host string, obj *orderedmap.OrderedMap) (err error) {
	id, ok := utils.ParseObjValueToString(obj, "id")
	if !ok {
		return errors.New("id not found in object")
	}
	actor, ok := utils.ParseObjValueToString(obj, "actor")
	if !ok {
		return errors.New("actor not found in object")
	}
	object, ok := utils.ParseObjValueToString(obj, "object")
	if !ok {
		return errors.New("object not found in object")
	}

	undoObj := activitypub.Undo(id+"#follows/undo", actor, obj)

	// endpoint
	inbox := object + "/inbox"
	// sign object
	req, err := signObjectByUser(undoObj, user, host, inbox)
	if err != nil {
		return err
	}
	_, err = fetchObjectWithRequest(req)
	return
}

func Reject(user, host string, obj *orderedmap.OrderedMap) (err error) {
	// id, ok := utils.ParseObjValueToString(obj, "id")
	// if !ok {
	// 	return errors.New("id not found in object")
	// }
	actor, ok := utils.ParseObjValueToString(obj, "actor")
	if !ok {
		return errors.New("actor not found in object")
	}
	// object, ok := utils.ParseObjValueToString(obj, "object")
	// if !ok {
	// 	return errors.New("object not found in object")
	// }

	rejectObj := activitypub.Reject(
		activitypub.APUserID(user, host)+"#follows/reject",
		activitypub.APUserID(user, host),
		obj,
	)

	// endpoint
	inbox := actor + "/inbox"
	// sign object
	req, err := signObjectByUser(rejectObj, user, host, inbox)
	if err != nil {
		return err
	}
	_, err = fetchObjectWithRequest(req)
	return
}

func Accept(user, host string, obj *orderedmap.OrderedMap) (err error) {
	// id, ok := utils.ParseObjValueToString(obj, "id")
	// if !ok {
	// 	return errors.New("id not found in object")
	// }
	actor, ok := utils.ParseObjValueToString(obj, "actor")
	if !ok {
		return errors.New("actor not found in object")
	}
	// object, ok := utils.ParseObjValueToString(obj, "object")
	// if !ok {
	// 	return errors.New("object not found in object")
	// }

	acceptObj := activitypub.Accept(
		activitypub.APUserID(user, host)+"#follows/accpet",
		activitypub.APUserID(user, host),
		obj,
	)

	// endpoint
	inbox := actor + "/inbox"
	// sign object
	req, err := signObjectByUser(acceptObj, user, host, inbox)
	if err != nil {
		return err
	}
	_, err = fetchObjectWithRequest(req)
	return
}
