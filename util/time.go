package util

import "time"

// GetCurrentTimestamp returns the current timestamp in miliseconds
func GetCurrentTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
