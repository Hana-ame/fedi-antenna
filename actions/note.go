package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Hana-ame/fedi-antenna/Tools/myfetch"
	"github.com/Hana-ame/fedi-antenna/actions/fetch"
	activitypub "github.com/Hana-ame/fedi-antenna/actions/model"
)

func CreateNote(note *activitypub.Note) error {
	if note == nil {
		log.Printf("nothing passed\n")
		return fmt.Errorf("nothing passed")
	}
	o := &activitypub.Create{
		Object: note,
	}
	o.Autofill()

	body, err := json.Marshal(o)
	if err != nil {
		return err
	}

	endpoint := "https://me.ns.ci/inbox"
	// endpoint := "https://relay.nya.one/inbox"
	// for _, endpoint := range utils.Endpoints {
	if resp, err := fetch.FetchWithSign(
		note.GetActor(),
		http.MethodPost, endpoint, nil, body,
	); err != nil {
		log.Println(err)
		return err
	} else {
		// debug only
		oo, err := myfetch.ResponseToObject(resp)
		log.Println(oo)
		log.Println(err)
		log.Println(resp.StatusCode)
	}
	return nil
}
