package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// https://api.lyhuilin.com/ping
// {"code":0,"message":"pong","data":"28.11.0"}
type pingPoneResponse struct {
	Code    int64  `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

// 自检openAPI服务是否正常运行
func CheckPingBaseURL(baseURL string) (retBool bool) {
	if !strings.HasPrefix(baseURL, "http") {
		return
	}
	apiURL := baseURL + "/ping"

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return
	}
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	// {"code":0,"message":"pong","data":"28.11.0"}
	var pingPong pingPoneResponse
	if err := json.Unmarshal(body, &pingPong); err != nil {
		return
	}
	if pingPong.Code > 0 || pingPong.Message != "pong" {
		return
	}
	if pingPong.Message == "pong" {
		return true
	}
	return true
}
