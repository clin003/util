package util

import (
	"io/ioutil"
	// "os"
	// "github.com/sirupsen/logrus"
	// "gitee.com/lyhuilin/log"
)

// ReadFile 读取文件
// 读取失败返回 nil
func ReadFile(path string) (retByte []byte, err error) {
	retByte, err = ioutil.ReadFile(path)
	if err != nil {
		// logrus.WithError(err).WithField("util", "ReadFile").Errorf("unable to read '%s'", path)
		// log.Errorf(err, "unable to read '%s'", path)
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
