package util

import (
	"fmt"
	"io/ioutil"
	"strings"

	// "gitee.com/lyhuilin/open_api/util"

	// "fmt"

	"net/http"
)

// 可还原的链接：
// https://u.jd.com/K7oUzZ
func GetRedirectUrlJD(str string, isClear bool) (retText string, err error) {
	req, err := http.NewRequest("GET", str, nil)
	if err != nil {
		return
	}
	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	httpClient := &http.Client{}
	// httpClient.Timeout = app.Timeout

	response, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	// fmt.Println("response.StatusCode", response.StatusCode)
	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	// strings.TrimSpace(s string)会返回一个string类型的slice，并将最前面和最后面的ASCII定义的空格去掉，中间的空格不会去掉，如果遇到了\0等其他字符会认为是非空格。
	// retBody := strings.TrimSpace(string(body))
	// 清除网页内容中的空格
	retBody := strings.Replace(string(body), " ", "", -1)
	retText = BetweenStr(retBody, "hrl='", "'")
	if retText != "" {
		retText, err = GetRedirectUrl(retText, false)
	}
	return
}
