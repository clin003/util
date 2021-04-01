package util

import (
	"bytes"
	"strings"
)

// import "strings"

// 取字符串中间 如果未找到，原样返回。
func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

// 取字符串中间，如果未找到返回空字符串
func BetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n < 0 {
		return ""

	}

	n += len(start)

	m := strings.Index(str[n:], end)

	if m < 0 {
		return ""
	}
	return str[n : n+m]
}

// 取字符串左侧
func LeftStr(str, posStr string, isContainsPosStr bool) string {
	n := strings.Index(str, posStr)
	if n < 0 {
		return ""
	}
	if isContainsPosStr {
		n += len(posStr)
	}

	// m := strings.Index(str[n:], end)

	// if m < 0 {
	// 	return ""
	// }
	return str[:n]
}

// 取字符串右侧
func RightStr(str, posStr string, isContainsPosStr bool) string {
	// 返回指定字符在此字符串中最后一次出现处的索引
	n := strings.LastIndex(str, posStr)
	if n < 0 {
		return ""
	}
	if !isContainsPosStr {
		n += len(posStr)
	}

	// m := strings.Index(str[n:], end)

	// if m < 0 {
	// 	return ""
	// }
	return str[n:]
}

//补全网址协议头
func UrlPatchHttp(url string) (ret string) {
	if strings.Index(url, "http") < 0 && url != "" {
		var buffer bytes.Buffer
		buffer.WriteString("https:")
		buffer.WriteString(url)
		return buffer.String()
	}
	return url
}

// //补全网址协议头
// func urlPatchHttp(url string) (ret string) {
// 	if strings.Index(url, "http") < 0 && url != "" {
// 		var buffer bytes.Buffer
// 		buffer.WriteString("https:")
// 		buffer.WriteString(url)
// 		return buffer.String()
// 	}
// 	return url
// }

// //取字符串中间
// func between(str, starting, ending string) string {
// 	s := strings.Index(str, starting)
// 	if s < 0 {
// 		return ""
// 	}
// 	s += len(starting)
// 	e := strings.Index(str[s:], ending)
// 	if e < 0 {
// 		return ""
// 	}
// 	return str[s : s+e]
// }

// func SubStrRange(s string, length int) string {
// 	var n, i int
// 	for i = range s {
// 		if n == length {
// 			break
// 		}

// 		n++
// 	}
// 	return s[:i]
// }
