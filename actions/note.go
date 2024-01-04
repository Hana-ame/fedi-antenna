package actions

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/activitypub/actions"
	activitypub "github.com/Hana-ame/fedi-antenna/activitypub/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

func CreateNote(note activitypub.Creatable) error {
	o := &activitypub.Create{
		Object: note,
	}
	o.Autofill()

	body, err := json.Marshal(o)
	if err != nil {
		return err
	}
	for _, endpoint := range utils.Endpoints {
		_, err := actions.FetchWithSign(
			note.GetActor(),
			http.MethodPost, endpoint, nil, body,
		)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
