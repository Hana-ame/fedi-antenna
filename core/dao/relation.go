package dao

import (
	"github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

// 这是啥
// 不想改表了，很麻烦。
// not tested.
func ReadFollowersByLocaluserID(id string) (sharedInboxes []string, err error) {
	// 1. query
	var relations []model.LocalRelation
	tx := db.Where("object = ?", id).Find(&relations)
	err = tx.Error
	if err != nil {
		return
	}

	// 2. select hosts.

	// 2.1. remove more than once
	var m map[string]any = make(map[string]any)
	for _, r := range relations {
		_, host := utils.ParseNameAndHost(r.Actor)
		m[host] = true
	}
	// 2.2. host to shared inbox
	var sharedInbox []model.FediStatus
	var hosts []string = make([]string, len(m))
	i := 0
	for k, _ := range m {
		hosts[i] = k
		i++
	}
	db.Find(&sharedInbox, hosts)

	// 3. move into []string
	sharedInboxes = make([]string, len(sharedInbox))
	for i, sharedinbox := range sharedInbox {
		sharedInboxes[i] = sharedinbox.SharedInbox
	}
	return
}
