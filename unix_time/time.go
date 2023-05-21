package unix_time

import (
	"time"
)

func UnixMilli() int64 {
	return time.Now().UnixMilli()
}

func UnixNano() int64 {
	return time.Now().UnixNano()
}