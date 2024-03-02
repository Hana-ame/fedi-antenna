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

func knownServers() map[string]struct{} {
	return map[string]struct{}{
		"o3o.ca":    struct{}{},
		"pawoo.net": struct{}{},
	}
}

func actorsServers(actor string) map[string]struct{} {
	dao.ReadFollowersByLocaluserID(actor)
}

func DoTask(actionType, dataType, primarykey, inbox, actor string) error {
	if inbox == "" {
		if actionType == model.TypeCreate && dataType == model.TypeNote {
			// query the visibility

		} else if actionType == model.TypeAnnounce {
			// query the visibility
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

	return nil
}
func AddTask(typ, primarykey, inbox, actor string) error {
	u, err := url.Parse(inbox)
	if err != nil {
		return err
	}
	task := &model.Task{
		AddedAt:    utils.NewTimestamp(false),
		Host:       u.Host,
		Inbox:      inbox,
		Type:       typ,
		ForeignKey: primarykey,

		Status: model.TaskPadding,
	}

	return nil
}

func exec(task *model.Task) error {

}

type Agent struct {
	q *queue.Queue
}
