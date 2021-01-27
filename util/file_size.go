package util

import "fmt"

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
)

// 转化文件大小到人工可读的字符串，以相应的单位显示
func FormatFSize(fsize int64) (result string) {
	if fsize > TB {
		result = fmt.Sprintf("%.2f TB", float64(fsize)/float64(TB))
	} else if fsize > GB {
		result = fmt.Sprintf("%.2f GB", float64(fsize)/float64(GB))
	} else if fsize > MB {
		result = fmt.Sprintf("%.2f MB", float64(fsize)/float64(MB))
	} else if fsize > KB {
		result = fmt.Sprintf("%.2f KB", float64(fsize)/float64(KB))
	} else {
		result = fmt.Sprintf("%d B", fsize)
	}
	return
}
