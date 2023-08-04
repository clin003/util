package util

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

const (
	urlReg = `((https|http|ftp|rtsp|mms)?:\/\/)[^\s]+`
)

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

func UrlRegMatchReplaceToHTML(str string) (ret string, err error) {
	ret = str
	tplParseText := str
	compile := regexp.MustCompile(urlReg)
	urlMap := make(map[string]string)
	submatch := compile.FindAllSubmatch([]byte(ret), -1)
	for index, m := range submatch {
		url := string(m[0])
		urlKey := "url" + fmt.Sprintf("%d", index)
		// fmt.Println("index", index, "url", url, "urlKey", urlKey)
		urlMap[urlKey] = url
		tplParseText = strings.ReplaceAll(tplParseText, url, fmt.Sprintf("<a href=\"{{ .%s }}\">{{ .%s }}</a>", urlKey, urlKey))
	}
	// fmt.Println("tplParseText", tplParseText)
	var result bytes.Buffer
	tpl, err := template.New("message").Parse(tplParseText)
	if err != nil {
		return ret, fmt.Errorf("template parsing failed: %w", err)
	}
	if err := tpl.Execute(&result, urlMap); err != nil {
		return ret, fmt.Errorf("template execution failed: %w", err)
	}
	retHtml := result.String()
	return retHtml, nil
}
