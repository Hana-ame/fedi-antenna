package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/actions/fetch"
	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
	"github.com/Hana-ame/fedi-antenna/core/convert"
	core "github.com/Hana-ame/fedi-antenna/core/model"
	"github.com/Hana-ame/fedi-antenna/core/utils"
)

func CreateNote(note *activitypub.Note, localnote *core.LocalNote, shouldRead bool) error {
	if note == nil {
		if localnote == nil {
			log.Printf("nothing passed\n")
			return fmt.Errorf("nothing passed")
		}
		note = convert.ToActivityPubNote(localnote)
	}
	o := &activitypub.Create{
		Object: note,
	}
	o.Autofill()

	body, err := json.Marshal(o)
	if err != nil {
		return err
	}
	for _, endpoint := range utils.Endpoints {
		_, err := fetch.FetchWithSign(
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
