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
		if isChinese.MatchString(vc) || unicode.IsLetter(vcr) {
		} else {
			textTmp = strings.ReplaceAll(textTmp, vc, "")
		}
	}
	// 清理变异字
	textTmp = strings.ReplaceAll(textTmp, "元", "")
	textTmp = strings.ReplaceAll(textTmp, "亓", "")
	textTmp = strings.ReplaceAll(textTmp, "劵", "卷")
	textTmp = strings.ReplaceAll(textTmp, "拎", "领")
	textTmp = strings.ReplaceAll(textTmp, "単", "单")
	textTmp = strings.ReplaceAll(textTmp, "吤", "价")
	textTmp = strings.ReplaceAll(textTmp, "菅", "营")
	textTmp = strings.ReplaceAll(textTmp, "咐", "付")
	textTmp = strings.ReplaceAll(textTmp, "歀", "款")
	textTmp = strings.ReplaceAll(textTmp, "啇", "商")

	textTmp = strings.ReplaceAll(textTmp, "便冝", "便宜")
	textTmp = strings.ReplaceAll(textTmp, "消售", "销售")
	textTmp = strings.ReplaceAll(textTmp, "復挃", "复制")
	textTmp = strings.ReplaceAll(textTmp, "津帖", "津贴")
	textTmp = strings.ReplaceAll(textTmp, "厍存", "库存")
	textTmp = strings.ReplaceAll(textTmp, "妙杀", "秒杀")
	textTmp = strings.ReplaceAll(textTmp, "欲售", "预售")
	textTmp = strings.ReplaceAll(textTmp, "钉金", "定金")
	textTmp = strings.ReplaceAll(textTmp, "订金", "定金")
	textTmp = strings.ReplaceAll(textTmp, "够物", "购物")
	textTmp = strings.ReplaceAll(textTmp, "优慧", "优惠")
	textTmp = strings.ReplaceAll(textTmp, "郵费", "运费")
	textTmp = strings.ReplaceAll(textTmp, "郵", "邮")
	textTmp = strings.ReplaceAll(textTmp, "航旗", "旗舰")

	textTmp = strings.ReplaceAll(textTmp, "活费", "话费")
	textTmp = strings.ReplaceAll(textTmp, "慢冲", "慢充")
	textTmp = strings.ReplaceAll(textTmp, "荭苞", "红包")
	textTmp = strings.ReplaceAll(textTmp, "虹包", "红包")
	textTmp = strings.ReplaceAll(textTmp, "兔費", "免费")

	textTmp = strings.ReplaceAll(textTmp, "亰", "京")
	textTmp = strings.ReplaceAll(textTmp, "荟", "会")
	textTmp = strings.ReplaceAll(textTmp, "桃包", "淘宝")
	textTmp = strings.ReplaceAll(textTmp, "喵抄", "猫超")
	textTmp = strings.ReplaceAll(textTmp, "喵", "猫")

	if len(textTmp) > 0 {
		retText = textTmp
		retHash = EncryptMd5(textTmp)
	}
	return
}
