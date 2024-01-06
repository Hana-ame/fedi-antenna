package utils

import (
	"sync/atomic"
	"time"
)

// 2023-08-22T10:23:08Z
func TimestampToRFC3339(ms int64) string {
	t := time.Unix(ms/1e6, 0).UTC()
	s := t.Format(time.RFC3339)
	return s
}

var timestamp atomic.Int64

// second * 1e6
func Now() (ms int64) {
	ms = time.Now().UnixNano() / 1e3
	if timestamp.Load() == ms {
		// time.Sleep(time.Microsecond)
		ms++
	}
	timestamp.Store(ms)
	return ms
}
