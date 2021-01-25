package qn_util

import "time"

func DateWithTimestamp(seconds int64, nanoSeconds int64) string {
	t := time.Unix(seconds, nanoSeconds)
	return t.String()
}

func Timestamp() int64 {
	return time.Now().Unix()
}