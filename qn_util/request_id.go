package qn_util

import (
	"encoding/base64"
	"fmt"
	"qshell/qn_error"
	"strconv"
	"time"
)

// 解析reqid， 打印人工可读的字符串
func ParseReqId(requestId string) (result string, err qn_error.IError) {
	decodedBytes, errD := base64.URLEncoding.DecodeString(requestId)
	if errD != nil || len(decodedBytes) < 4 {
		err = qn_error.NewCryptError(errD.Error())
		return
	}

	newBytes := decodedBytes[4:]
	newBytesLen := len(newBytes)
	newStr := ""
	for i := newBytesLen - 1; i >= 0; i-- {
		newStr += fmt.Sprintf("%02X", newBytes[i])
	}
	unixNano, errP := strconv.ParseInt(newStr, 16, 64)
	if errP != nil {
		err = qn_error.NewCryptError(errP.Error())
		return
	}

	dstDate := time.Unix(0, unixNano)
	result = fmt.Sprintf("%04d-%02d-%02d %02d:%02d", dstDate.Year(), dstDate.Month(), dstDate.Day(),
		dstDate.Hour(), dstDate.Minute())
	return
}