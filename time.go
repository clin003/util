package util

import (
	"time"
)

// 13位时间戳 UnixMilli
func TimeStampNow() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// fmt.Println(time.Now().Unix())
// fmt.Println(time.Now().UnixNano() / 1e6)
// 13位时间戳 UnixMilli
// isSecond=true 返回10位时间戳。isSecond=false 返回13位时间戳
func TimeStampNowEx(isSecond bool) int64 {
	if isSecond {
		return time.Now().Unix()
	} else {
		return time.Now().UnixNano() / int64(time.Millisecond)
	}
}
