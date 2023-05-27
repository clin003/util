package main

import (
	"strings"
	"sync"
)

var lockCacheCrosscheckKey sync.Mutex
var cacheCrosscheckKey []string

// 检查 key 是否存在 true 真 存在，false 不存在/或key为空，不存在存储
func CrosscheckKey(key string, maxKey int) (retBool bool) {
	if len(key) <= 0 {
		return false
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

// 检查 msgmsgSignature 是否已记录 true 真 存在，false 不存在，不存在存储
func MsgSignatureCheck(msgmsgSignature string, maxKey int) (retBool bool) {
	return CrosscheckKey(msgmsgSignature, maxKey)
}
func MsgTextCleanerCheck(msgText string, maxKey int) (retBool bool) {
	_, msgCleanerHash := GetTextCleaner(msgText)
	return CrosscheckKey(msgCleanerHash, maxKey)
}

func MsgSignatureCheckEx(msgmsgSignature, msgText string, maxKey int) (retBool bool) {
	return MsgSignatureCheck(msgmsgSignature, maxKey) || MsgTextCleanerCheck(msgText, maxKey)
}
