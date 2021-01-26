package qn_util

import (
	"encoding/base64"
	"qshell/qn_error"
)

func Base64Encode(value string, urlSafe bool) (result string, err qn_error.IError) {
	var (
		dataDecoded []byte
		errD error
	)

	if urlSafe {
		dataDecoded, errD = base64.URLEncoding.DecodeString(value)
	} else {
		dataDecoded, errD = base64.StdEncoding.DecodeString(value)
	}

	if errD != nil {
		err = qn_error.NewCryptError(errD.Error())
	} else {
		result = string(dataDecoded)
	}
	return
}

func Base64Decode(value string, urlSafe bool) string {

	var result string

	if urlSafe {
		result = base64.URLEncoding.EncodeToString([]byte(value))
	} else {
		result = base64.StdEncoding.EncodeToString([]byte(value))
	}

	return result
}