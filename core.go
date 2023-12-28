package core

import (
	"encoding/json"

	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/core/dao"
)

func init() {
	initHost()
}

// handlers

// public interface
func Inbox(b []byte, user string) error {
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	if v, ok := m["type"]; v == "Follow" && ok {
		var o *activitypub.Follow
		if err := json.Unmarshal(b, &o); err != nil {
			return err
		}
		return Follow(o)
	} else if v, ok := m["type"]; v == "Block" && ok {
		var o *activitypub.Block
		if err := json.Unmarshal(b, &o); err != nil {
			return err
		}
		return Block(o)
	} else if v, ok := m["type"]; v == "Undo" && ok {
		var o *activitypub.Undo
		if err := json.Unmarshal(b, &o); err != nil {
			return err
		}
		return Undo(o)
	}

	return nil
}

func Follow(o *activitypub.Follow) error {
	err := dao.Create(o)
	return err
}

func Undo(o *activitypub.Undo) error {
	err := dao.Delete(o.Object)
	return err
}

func Block(o *activitypub.Block) error {
	err := dao.Create(o)
	return err
}
