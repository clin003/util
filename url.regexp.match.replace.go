package util

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	urlReg = `((https|http|ftp|rtsp|mms)?:\/\/)[^\s]+`
)

// func strRegMatch(str, regStr string) (ret string) {
// 	ret = str
// 	compile := regexp.MustCompile(regStr)
// 	valueMap := make(map[string]string)
// 	submatch := compile.FindAllSubmatch([]byte(ret), -1)
// 	for _, m := range submatch {
// 		value := string(m[0])
// 		valueMap[value] = fmt.Sprintf("[value=%s]", value)
// 	}
// 	for k, v := range valueMap {
// 		ret = strings.ReplaceAll(ret, k, v)
// 		fmt.Printf(">>%s\n", ret)
// 	}
// 	return
// }

func UrlRegMatchReplace(str string) (ret string) {
	ret = str
	compile := regexp.MustCompile(urlReg)
	urlMap := make(map[string]string)
	submatch := compile.FindAllSubmatch([]byte(ret), -1)
	for _, m := range submatch {
		url := string(m[0])
		urlMap[url] = fmt.Sprintf("<a href=\"%s\">%s</a>", url, url)
	}
	// fmt.Println(urlMap)
	for k, v := range urlMap {
		ret = strings.ReplaceAll(ret, k, v)
	}
	// fmt.Println(ret)
	return
}

func UrlRegMatchReplaceToMarkdown(str string) (ret string) {
	ret = str
	compile := regexp.MustCompile(urlReg)
	urlMap := make(map[string]string)
	submatch := compile.FindAllSubmatch([]byte(ret), -1)
	for _, m := range submatch {
		url := string(m[0])
		urlMap[url] = fmt.Sprintf("[%s](%s)", url, url)
	}
	// fmt.Println(urlMap)
	for k, v := range urlMap {
		ret = strings.ReplaceAll(ret, k, v)
	}
	// fmt.Println(ret)
	return
}
