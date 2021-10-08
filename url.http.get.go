package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetUrlToString(url string) (retText string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode == 200 {
		retText = string(body)
		return
	} else {
		err = fmt.Errorf("response status code err:%d", resp.StatusCode)
	}
	return
}
func GetUrlToStringEx(url string, timeOut time.Duration) (retText string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	client := &http.Client{Timeout: timeOut}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode == 200 {
		retText = string(body)
		// retByte = body
		return
	} else {
		err = fmt.Errorf("response status code err:%d", resp.StatusCode)
	}
	return
}
func GetUrlToByte(url string) (retByte []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode == 200 {
		// retText = string(body)
		retByte = body
		return
	} else {
		err = fmt.Errorf("response status code err:%d", resp.StatusCode)
	}
	return
}
func GetUrlToByteEx(url string, timeOut time.Duration) (retByte []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	client := &http.Client{Timeout: timeOut}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode == 200 {
		// retText = string(body)
		retByte = body
		return
	} else {
		err = fmt.Errorf("response status code err:%d", resp.StatusCode)
	}
	return
}

func GetUrlWithCookieToString(url, cookie, host, referer, userAgent string) (retText string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	if len(cookie) > 0 {
		req.Header.Set("Cookie", cookie)
	}
	if len(host) > 0 {
		req.Header.Set("Host", host)
	}
	if len(referer) > 0 {
		req.Header.Set("Referer", referer)
	}
	if len(userAgent) > 0 {
		req.Header.Set("User-Agent", userAgent)
	}

	client := &http.Client{Timeout: time.Second * 30}
	resp, err := client.Do(req)
	if err != nil {
		// log.Fatal("Error reading response. ", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode == 200 {
		retText = string(body)
		return
	} else {
		err = fmt.Errorf("response status code err:%d", resp.StatusCode)
	}
	return
}

func GetUrlWithCookieToStringAndSaveFile(url, cookie, host, referer, userAgent, filePath string) (retText string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	if len(cookie) > 0 {
		req.Header.Set("Cookie", cookie)
	}
	if len(host) > 0 {
		req.Header.Set("Host", host)
	}
	if len(referer) > 0 {
		req.Header.Set("Referer", referer)
	}
	if len(userAgent) > 0 {
		req.Header.Set("User-Agent", userAgent)
	}

	client := &http.Client{Timeout: time.Second * 30}
	resp, err := client.Do(req)
	if err != nil {
		// log.Fatal("Error reading response. ", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode == 200 {
		retText = string(body)
		if len(filePath) > 0 {
			err = ioutil.WriteFile(filePath, body, 0666)
			// if err != nil {
			// 	// log.Errorf(err, "ioutil.WriteFile err")
			// }
		}
		return
	} else {
		err = fmt.Errorf("response status code err:%d", resp.StatusCode)
	}
	return
}
