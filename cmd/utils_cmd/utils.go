package utils_cmd

import (
	"qshell/cmd/param_cmd"
)

func LoadCMD(root param_cmd.IParamCMD) {
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