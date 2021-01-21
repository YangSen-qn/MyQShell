package qn_util

import (
	"io"
	"os"
	"qshell/qn_error"
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

func CreateFile(filename string, defaultData string) qn_error.IError {
	f, err := os.Create(filename)
	if err != nil{
		return qn_error.NewInvalidFilePathError(err.Error())
	}
	defer f.Close()

	_, err = io.WriteString(f, defaultData)
	if err != nil{
		return qn_error.NewInvalidIOError(err.Error())
	}

	return nil
}