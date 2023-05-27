package utils

import (
	"fmt"
	"regexp"
)

// const (
// 	urlReg = `((https|http|ftp|rtsp|mms)?:\/\/)[^\s]+` //url正则
// )

// 消息中链接识别
func UrlTextToMap(urlText string) (ret map[string]string, err error) {
	urlReg := `((https|http|ftp|rtsp|mms)?:\/\/)[^\s]+` //url正则
	//链接识别并转链
	compile := regexp.MustCompile(urlReg)
	urlMap := make(map[string]string)
	submatchUrl := compile.FindAllSubmatch([]byte(urlText), -1)
	for _, m := range submatchUrl {
		value := string(m[0])
		urlMap[value] = value
	} //链接转链
	if len(urlMap) <= 0 {
		err = fmt.Errorf("未识别到链接")
	}
	return urlMap, err
}
