package actions

import (
	"net/url"
	"sync"

	"github.com/Hana-ame/fedi-antenna/actions/model"
	"github.com/Hana-ame/fedi-antenna/actions/model/queue"
	"github.com/Hana-ame/fedi-antenna/core/dao"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

var mu sync.RWMutex
var agents = make(map[string]*Agent)

func Host2PublicInbox(host string) (inbox string) {
	return "https://" + host + "/inbox"
}

// todo
// return known servers
func knownServers() map[string]struct{} {
	return map[string]struct{}{
		"o3o.ca":    struct{}{},
		"pawoo.net": struct{}{},
	}
}

// todo
// return the followers
func followersServers(actor string) map[string]struct{} {
	dao.ReadFollowersByLocaluserID(actor)
	return map[string]struct{}{}
}

// todo
// return the inboxes should send
func inboxes(actionType, dataType, primarykey, inbox, actor string) (map[string]struct{}, error) {
	if actionType == model.TypeCreate && dataType == model.TypeNote {
		// query the visibility

	} else if actionType == model.TypeAnnounce {
		// query the visibility
	}
	return knownServers(), nil
}

func DoTask(actionType, dataType, primarykey, inbox, actor string) error {
	// when not specified inbox,
	if inbox == "" {
		inboxSet, err := inboxes(actionType, dataType, primarykey, inbox, actor)
		if err != nil {
			return err
		}
		for i := range inboxSet {
			DoTask(actionType, dataType, primarykey, i, actor)
		}
	}

	u, err := url.Parse(inbox)
	if err != nil {
		return err
	}
	task := &model.Task{
		AddedAt:    utils.NewTimestamp(false),
		Host:       u.Host,
		Inbox:      inbox,
		ActionType: actionType,
		DataType:   dataType,
		ForeignKey: primarykey,

		Status: model.TaskPadding,
	}
	tx := dao.Begin()
	if err := dao.Create(tx, task); err != nil {
		tx.Rollback()
		return err
	}

	return exec(task)
}

// todo
// should do the tasks here,
func exec(task *model.Task) error {
	return nil
}

type Agent struct {
	q *queue.Queue
}
