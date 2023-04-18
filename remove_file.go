package util

import (
	"fmt"
	"os"
	"strings"
)

func RemoveFile(filePath string) error {
	if len(filePath) <= 0 {
		return fmt.Errorf("filePath is Null")
	}
	if !strings.HasSuffix(filePath, ".mp4") {
		return fmt.Errorf("file is not mp4")
	}
	err := os.Remove(filePath)
	if err != nil {
		err = fmt.Errorf("remove(%s): %v", filePath, err)
		return err
	}
	fmt.Println("Remove File:", filePath)
	return nil
}
