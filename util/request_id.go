package util

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
)

// 解析reqId， 打印人工可读的字符串
func ParseReqId(requestId string) (result string, err error) {
	decodedBytes, err := base64.URLEncoding.DecodeString(requestId)
	if err != nil || len(decodedBytes) < 4 {
		return
	}

	newBytes := decodedBytes[4:]
	newBytesLen := len(newBytes)
	newStr := ""
	for i := newBytesLen - 1; i >= 0; i-- {
		newStr += fmt.Sprintf("%02X", newBytes[i])
	}
	unixNano, err := strconv.ParseInt(newStr, 16, 64)
	if err != nil {
		return
	}

	dstDate := time.Unix(0, unixNano)
	result = fmt.Sprintf("%04d-%02d-%02d %02d:%02d", dstDate.Year(), dstDate.Month(), dstDate.Day(),
		dstDate.Hour(), dstDate.Minute())
	return
}