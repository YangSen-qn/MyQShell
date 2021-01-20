package utils

import (
	"os"
	"qshell/qn_shell_error"
)

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func CreateFile(filename string) qn_shell_error.IQShellError {
	f, err := os.Create(filename)
	if err != nil{
		return qn_shell_error.NewInvalidFilePathError(err.Error())
	}
	f.Close()
	return nil
}