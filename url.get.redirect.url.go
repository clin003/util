package util

import (
	// "fmt"
	"time"

	"net/http"
)

// 可还原的链接：
// 淘宝 https://s.click.taobao.com/rioMXxu  二合一链接
// 苏宁 https://sugs.suning.com/HpGbzgk6
// https://sugs.suning.com/HpCPkl3B
// 唯品会 https://t.vip.com/MhW6RosYXqA
func GetRedirectUrl(str string, isClear bool) (string, error) {
	resp, err := http.Head(str)
	if err != nil {
		return str, err
	}
	defer resp.Body.Close()

	if isClear == false {
		newUrl := resp.Request.URL.String()
		return newUrl, nil
	}
	// fmt.Printf(">>%s\n", resp.Request.URL)
	newUrl := resp.Request.URL.Scheme + "://" + resp.Request.URL.Host + resp.Request.URL.Path
	return newUrl, nil
}

// 可还原的链接：
// 淘宝 https://s.click.taobao.com/rioMXxu  二合一链接
// 苏宁 https://sugs.suning.com/HpGbzgk6
// https://sugs.suning.com/HpCPkl3B
// 唯品会 https://t.vip.com/MhW6RosYXqA
// 拼多多 https://p.pinduoduo.com/IE9vjwqh
// 考拉 https://lu.kaola.com/gKGdm
func GetRedirectUrlEx(str string, isClear bool) (retText string, err error) {
	req, err := http.NewRequest("GET", str, nil)
	if err != nil {
		return
	}
	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	httpClient := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: 30 * time.Second,
	}

	response, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	// fmt.Println("response.StatusCode", response.StatusCode)
	if response.StatusCode == 200 {
		// err = fmt.Errorf("请求错误:%d", response.StatusCode)
		retText = response.Request.URL.String()
		return
	}
	retText = response.Header.Get("Location")
	// fmt.Println("response.Header", response.Header)
	// retText = location
	return
}
