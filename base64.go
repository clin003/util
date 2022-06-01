package util

import (
	"encoding/base64"
)

func Base64StdEncodingEncodeToString(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}
func Base64StdEncodingDecode(input string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(input)
}
