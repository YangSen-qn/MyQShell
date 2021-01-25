package utils_cmd

import (
	"qshell/cmd/param_cmd"
)

func LoadCMD(root param_cmd.IParamCMD) {
	loadDateCMD(root)
	loadTimestampCMD(root)
}