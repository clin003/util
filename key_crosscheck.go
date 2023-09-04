package util

import (
	"strings"
	"sync"
)

var lockCacheCrosscheckKey sync.Mutex
var cacheCrosscheckKey []string

// 检查 key 是否存在 true 真 存在，false 不存在/或key为空，不存在存储
func CrosscheckKey(key string, maxKey int) (retBool bool) {
	if len(key) <= 0 || strings.EqualFold(key, "k1") || strings.EqualFold(key, "k2") {
		return false
	}
	if maxKey > 10000 || maxKey <= 0 {
		maxKey = 10000
	}
	lockCacheCrosscheckKey.Lock()
	defer lockCacheCrosscheckKey.Unlock()

	for _, v := range cacheCrosscheckKey {
		if strings.EqualFold(v, key) {
			return true
		}
	}
	sliceTmp := make([]string, 1)
	sliceTmp[0] = key
	if len(cacheCrosscheckKey) > maxKey {
		cacheCrosscheckKey = append(sliceTmp, cacheCrosscheckKey[0:len(cacheCrosscheckKey)-1]...)
	} else {
		cacheCrosscheckKey = append(sliceTmp, cacheCrosscheckKey[0:]...)
	}

	return false
}
func CrosscheckKeyWithBase(keyBase, key string, maxKey int) (retBool bool) {
	if len(key) <= 0 || strings.EqualFold(key, keyBase) {
		return false
	}
	return CrosscheckKey(key, maxKey)
}

// 检查 msgmsgSignature 是否已记录 true 真 存在，false 不存在，不存在存储
func MsgSignatureCheck(msgmsgSignature string, maxKey int) (retBool bool) {
	return CrosscheckKeyWithBase("k1", "k1"+msgmsgSignature, maxKey)
}
func MsgTextCleanerCheck(msgText string, maxKey int) (retBool bool) {
	_, msgCleanerHash := GetTextCleaner(msgText)
	return CrosscheckKeyWithBase("k2", "k2"+msgCleanerHash, maxKey)
}

// 检查 msgmsgSignature 是否已记录 true 真 存在，false 不存在，不存在存储
func MsgSignatureCheckEx(msgmsgSignature, msgText string, maxKey int) (retBool bool) {
	return MsgSignatureCheck(msgmsgSignature, maxKey) || MsgTextCleanerCheck(msgText, maxKey)
}
