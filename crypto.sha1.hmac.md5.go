package util

import (
	"crypto/md5"
	"crypto/sha1"

	"crypto/hmac"
	"encoding/hex"
	"fmt"
)

// hash plaintext with SHA-1
func EncryptSha1(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}

// hash plaintext with md5
func EncryptMd5(source string) (cryptext string) {
	if len(source) <= 0 {
		return ""
	}
	// 使用MD5加密
	signBytes := md5.Sum([]byte(source))
	// 把二进制转化为大写的十六进制
	cryptext = hex.EncodeToString(signBytes[:])
	return
}

// hash plaintext with md5
func EncryptMd5Byte(source []byte) (cryptext string) {
	// 使用MD5加密
	signBytes := md5.Sum(source)
	// 把二进制转化为大写的十六进制
	cryptext = hex.EncodeToString(signBytes[:])
	return
}

// https://www.jianshu.com/p/02da10ca45d1
// hmac-md5加密
func EncryptHmacmd5(src, key string) (cryptext string) {
	// m:=md5.New()
	m := hmac.New(md5.New, []byte(key))
	m.Write([]byte(src))
	cryptext = hex.EncodeToString(m.Sum(nil))
	return
}
