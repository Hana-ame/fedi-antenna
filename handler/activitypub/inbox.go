package handler

import (
	"encoding/json"

	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
)

// handlers

// public interface
func Inbox(b []byte, user string) error {
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	if v, ok := m["type"]; v == "Create" && ok {
		var o *activitypub.Create
		if err := json.Unmarshal(b, &o); err != nil {
			return err
			return Create(o)
		}
	} else if v, ok := m["type"]; v == "Follow" && ok {
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
