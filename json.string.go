package util

import (
	"bytes"
	"encoding/json"
)

func Marshal(obj interface{}) string {
	var bytes, err = json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// json编码后，是否对HTML内容进行转义处理。
func MarshalEx(obj interface{}, isEscapeHTML bool) string {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(isEscapeHTML)
	jsonEncoder.Encode(obj)
	return bf.String()
}
