package utils

import (
	"qshell/cmd/param"
)

func LoadCMD(root param.IParamCMD) {
	loadDateCMD(root)
	loadTimestampCMD(root)
	loadUrlCMD(root)
	loadEtagCMD(root)
	loadFileSizeCMD(root)
	loadRPCCMD(root)
	loadParseRequestIdCMD(root)
	loadBase64CMD(root)
	loadZipCMD(root)
}