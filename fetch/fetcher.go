package fetch

import (
	"encoding/json"
	"log"
	"time"
)

type sharedInbox map[string]*Fetcher

var PeerSites sharedInbox = make(sharedInbox)

type Fetcher struct {
	Host        string
	SharedInbox string
}

func NewFetcher(host, sharedInbox string) *Fetcher {
	o := &Fetcher{
		Host:        host,
		SharedInbox: sharedInbox,
	}
	PeerSites[host] = o
	return o
}

func (f *Fetcher) WithSign(method string, o any, inbox string, localuserID string) (err error) {
	if inbox == "" {
		inbox = f.SharedInbox
	}
	for i := 0; i < 5; i++ {
		err = f.withSign(method, o, inbox, localuserID)
		if err == nil {
			break
		}
		for j := i; j >= 0; j++ {
			time.Sleep(60 * time.Second)
		}
	}
	return
}

func (f *Fetcher) withSign(method string, o any, inbox string, localuserID string) (err error) {

	body, err := json.Marshal(o)
	if err != nil {
		return err
	}

	FetchWithSign(
		localuserID,
		method, inbox, nil, body,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
