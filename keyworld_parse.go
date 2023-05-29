package util

import (
	"strings"
)

// 解析 k|k1 到数组
func KeyworldListParse(keyworldListStr string) []string {
	if len(keyworldListStr) <= 0 {
		return nil
	}

	retArr := make([]string, 0)
	listText := keyworldListStr
	list := strings.Split(listText, "|")
	for _, v := range list {
		vc := v
		if len(vc) > 0 {
			retArr = append(retArr, vc)
		}
	}
	return retArr
}

//解析 k>>kc|k1>>kc1 到 map
func KeyworldListParseToMap(keyworldText string) map[string]string {
	keyworldList := KeyworldListParse(keyworldText)
	retMap := make(map[string]string, 0)
	for _, v := range keyworldList {
		if len(v) > 0 {
			if strings.Contains(v, ">>") {
				if vk, vv, isOk := strings.Cut(v, ">>"); isOk {
					if len(vk) > 0 {
						retMap[vk] = vv
					}
				}
			} else {
				retMap[v] = ""
			}
		}
	}
	return retMap
}
