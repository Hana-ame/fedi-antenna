package utils

import "time"

// 2023-08-22T10:23:08Z
func MicroSecondToRFC3339(ms int64) string {
	t := time.Unix(ms/1e3, ms%1e3*1e6).UTC()
	s := t.Format(time.RFC3339)
	return s
}

//
func Now() (ms int64) {
	ms = time.Now().UnixMilli()
	return ms
}
