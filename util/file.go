package util

import (
	"io"
	"os"
)

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func IsFileEmpty(filename string) bool {
	if !Exist(filename) {
		return true
	}

	fileInfo, _ := os.Stat(filename)
	if fileInfo != nil && fileInfo.Size() > 0 {
		return false
	} else {
		return true
	}
}

func CreateFile(filename string, defaultData string) error {
	f, err := os.Create(filename)
	if err != nil{
		return err
	}
	defer f.Close()

	_, err = io.WriteString(f, defaultData)
	if err != nil{
		return err
	}

	return nil
}