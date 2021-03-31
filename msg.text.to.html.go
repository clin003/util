package util

import (
	"strings"
)

func MsgTextToHtml(msg string) (retText string, err error) {
	contentEx := UrlRegMatchReplace(msg)
	contentEx = strings.Replace(contentEx, "\n", "<br>", -1)
	if len(contentEx) > 0 {
		retText = contentEx
	}
	return
}
