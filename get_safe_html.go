package util

import (
	"bytes"
	"html/template"
)

// 获取 text 的 安全html格式内容（转义html）
func GetSafeHtml(text string) (string, error) {
	type HtmlContent struct {
		Content string
	}
	var htmlContent HtmlContent
	htmlContent.Content = text
	var result bytes.Buffer
	newTpl := func() (*template.Template, error) {
		tplParseText := "{{ .Content }}"
		return template.New("message").Parse(tplParseText)
	}

	tpl, err := newTpl()
	if err != nil {
		return "", err
	}
	if err := tpl.Execute(&result, htmlContent); err != nil {
		return "", err
	}
	return result.String(), nil
}
