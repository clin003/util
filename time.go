package util

import (
	"time"
)

// 13位时间戳 UnixMilli
func TimeStampNow() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
