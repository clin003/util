package util

import (
	"io/ioutil"
)

// ReadFile 读取文件
// 读取失败返回 nil
func ReadFile(path string) (retByte []byte, err error) {
	retByte, err = ioutil.ReadFile(path)
	if err != nil {
		return //nil, err
	}
	return //bytes
}

func WriteFile(filePath string, fileBody []byte) error {
	// err =ioutil.WriteFile(filePath, fileBody, 0666)
	return ioutil.WriteFile(filePath, fileBody, 0666)
}

// FileExist 判断文件是否存在
func IsFileExist(path string) (bool, error) {
	return IsExists(path)
}
