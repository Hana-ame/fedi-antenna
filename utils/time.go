package utils

import "time"

// 2023-08-22T10:23:08Z
func TimestampToRFC3339(ms int64) string {
	t := time.Unix(ms/1e3, 0).UTC()
	s := t.Format(time.RFC3339)
	return s
}

// second * 1e3
func Now() (ms int64) {
	ms = time.Now().UnixNano() / 1e6
	return ms
}
