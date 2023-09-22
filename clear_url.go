package util

import (
	"regexp"
)

// 清除文本中 url
func ClearUrl(text string) string {
	reg := regexp.MustCompile(`((https|http|ftp|rtsp|mms)?:\/\/)[^\s]+`)
	return reg.ReplaceAllString(text, "")
}
