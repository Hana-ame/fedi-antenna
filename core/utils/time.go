package utils

import (
	"sync"
	"time"
)

// 2023-08-22T10:23:08Z
func TimestampToRFC3339(us int64) string {
	t := time.Unix(us/1e6, 0).UTC()
	s := t.Format(time.RFC3339)
	return s
}

var mutex sync.Mutex
var timestamp int64

// second * 1e6
func Timestamp(fast bool) (now int64) {
	now = time.Now().UnixNano() / 1e3
	if fast {
		return now

	}
	mutex.Lock()
	if timestamp == now {
		now++
	}
	timestamp = now
	defer mutex.Unlock()

	return now
}
