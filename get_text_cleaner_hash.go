package util

import (
	"regexp"
	"strings"
	"unicode"
)

func GetTextCleaner(textStr string) (retText, retHash string) {
	textTmp := textStr
	urlMap, _ := UrlTextToMap(textTmp)
	for _, v := range urlMap {
		vc := v
		textTmp = strings.ReplaceAll(textTmp, vc, "")
	}

	isChinese := regexp.MustCompile("^[\u4e00-\u9fa5]") //我们要匹配中文的匹配规则
	//使用MatchString来将要匹配的字符串传到匹配规则中
	// unicode包的IsLetter方法判断字符是不是字母
	for _, v := range textTmp {
		vcr := v
		vc := string(v)
		// switch {
		// case isChinese.MatchString(vc):
		// case unicode.IsLetter(vcr):
		// default:
		// 	textTmp = strings.ReplaceAll(textTmp, vc, "")
		// }
		if isChinese.MatchString(vc) || unicode.IsLetter(vcr) {
		} else {
			textTmp = strings.ReplaceAll(textTmp, vc, "")
		}
	}
	textTmp = strings.ReplaceAll(textTmp, "元", "")
	textTmp = strings.ReplaceAll(textTmp, "亓", "")

	if len(textTmp) > 0 {
		retText = textTmp
		retHash = EncryptMd5(textTmp)
	}
	return
}
