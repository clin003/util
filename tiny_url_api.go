package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type tinyURLAPIResponse struct {
	Code int64 `json:"code"`
	Data struct {
		ShortURL string `json:"short_url"`
		URL      string `json:"url"`
	} `json:"data"`
	Message string `json:"message"`
}

// 长链接转短链接
// http://127.0.0.1:8089/tinyURL?url=http://www.lyhuilin.com
// {"code":0,"message":"OK","data":{"url":"http://www.lyhuilin.com","short_url":"https://u.nav.xin/4i"}}
func TinyURLAPI(langURL, apiBaseURL string, timeDuration time.Duration, isJson bool) (retText string, err error) {
	apiBaseSite := apiBaseURL

	p := url.Values{}
	p.Add("url", langURL)

	routerUrl := fmt.Sprintf("%s/tinyURL?%s", apiBaseSite, p.Encode())
	if !strings.HasPrefix(routerUrl, "http") {
		err = fmt.Errorf("apiBaseURL:%s", apiBaseURL)
		return
	}

	req, err := http.NewRequest("GET", routerUrl, nil)
	if err != nil {
		return
	}

	httpClient := &http.Client{Timeout: timeDuration}
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if isJson {
		retText = string(body)
		return
	}

	var retJson tinyURLAPIResponse
	err = json.Unmarshal(body, &retJson)
	if err != nil {
		return
	}

	if retJson.Code > 0 {
		err = fmt.Errorf("error:%s", string(body))
		return
	}
	if len(retJson.Data.ShortURL) > 0 {
		retText = retJson.Data.ShortURL
	} else {
		err = fmt.Errorf("error:%s", string(body))
	}
	return
}
