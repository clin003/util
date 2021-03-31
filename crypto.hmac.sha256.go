package util

import (
	"crypto/hmac"
	// "crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func EncryptHmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
