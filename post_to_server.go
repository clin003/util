package util

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

func PostToServer(serverApiUrl, serverToken string, postData interface{}) (ret string, err error) {
	jsonParams, err := json.Marshal(postData)
	if err != nil {
		return
	}
	jsonString := string(jsonParams)
	payload := strings.NewReader(jsonString)

	apiUrl := serverApiUrl
	req, err := http.NewRequest("POST", apiUrl, payload)
	if err != nil {
		return
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("apiclient", serverToken) //接收 c.Request.Header["Apiclient"]
	// req.Header.Add("User-Agent", "apisender")
	client := http.Client{
		Timeout: time.Second * 30,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	// fmt.Println(res)
	ret = string(body)
	return
}
