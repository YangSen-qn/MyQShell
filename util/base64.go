package util

import (
	"encoding/base64"
)

func Base64Encode(value string, urlSafe bool) (result string, err error) {
	var dataDecoded []byte

	if urlSafe {
		dataDecoded, err = base64.URLEncoding.DecodeString(value)
	} else {
		dataDecoded, err = base64.StdEncoding.DecodeString(value)
	}

	if err == nil {
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