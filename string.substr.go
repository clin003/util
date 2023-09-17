package util

import (
	"unicode/utf8"
)

// 截取字符串
func SubStrDecodeRuneInString(s string, length int) string {
	var size, n int
	for i := 0; i < length && n < len(s); i++ {
		_, size = utf8.DecodeRuneInString(s[n:])
		n += size
	}

	return s[:n]
}
func SubStrRange(s string, length int) string {
	var n, i int
	for i = range s {
		if n == length {
			break
		}

		n++
	}

	return s[:i]
}

// 截图字符串前N个字符内容
func SubStrRuneN(textStr string, n int) (retText string) {
	if len(textStr) <= 0 {
		return ""
	}
	retText = textStr
	textR := []rune(retText)
	if len(textR) > n {
		retText = string(textR[:n-1])
	}
	retText = strings.TrimSpace(retText)
	return retText
}
