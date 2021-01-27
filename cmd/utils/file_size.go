package utils

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param"
	"qshell/qn_error"
	"qshell/util"
	"strconv"
)

const (
	fileSizeEmptyValue = -1
)

type fileSizeCMD struct {
	*param.ParamCMD

	fileSize int64
}

func (cmd *fileSizeCMD) Prepare(context *common.QShellContext) error {
	fileSize, err := strconv.ParseInt(cmd.GetFirstArg(), 10, 64)
	if err == nil {
		cmd.fileSize = fileSize
	}
	return nil
}

func (cmd *fileSizeCMD) Check(context *common.QShellContext) error {
	if cmd.fileSize == fileSizeEmptyValue {
		return qn_error.NewInvalidUserParamError("fize size of Byte should not empty")
	} else {
		return nil
	}
}

func (cmd *fileSizeCMD) Execute(context *common.QShellContext) error {

	fileSize := util.FormatFSize(cmd.fileSize)
	output_utils.OutputResultWithString(cmd, fileSize)
	return nil
}

func loadFileSizeCMD(root param.IParamCMD) param.IParamCMD {
	cmd := &fileSizeCMD{
		fileSize: fileSizeEmptyValue,
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		PrepareFunction: cmd.Prepare,
		CheckFunction:   cmd.Check,
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "file-size",
		Short: "get sourceFile size from Byte",
		Long:  "",
	})

	root.AddCMD(cmd)
	return cmd
}
