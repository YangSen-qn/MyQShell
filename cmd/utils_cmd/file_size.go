package utils_cmd

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param_cmd"
	"qshell/qn_error"
	"qshell/qn_util"
	"strconv"
)

const (
	fileSizeEmptyValue = -1
)

type fileSizeCMD struct {
	*param_cmd.ParamCMD

	fileSize int64
}

func (cmd *fileSizeCMD) Prepare(context *common.QShellContext) qn_error.IError {
	fileSize, err := strconv.ParseInt(cmd.GetFirstArg(), 10, 64)
	if err == nil {
		cmd.fileSize = fileSize
	}
	return nil
}

func (cmd *fileSizeCMD) Check(context *common.QShellContext) qn_error.IError {
	if cmd.fileSize == fileSizeEmptyValue {
		return qn_error.NewInvalidUserParamError("fize size of Byte should not empty")
	} else {
		return nil
	}
}

func (cmd *fileSizeCMD) Execute(context *common.QShellContext) qn_error.IError {

	fileSize := qn_util.FormatFsize(cmd.fileSize)
	output_utils.OutputResultWithString(cmd, fileSize)
	return nil
}

func loadFileSizeCMD(root param_cmd.IParamCMD) param_cmd.IParamCMD {
	cmd := &fileSizeCMD{
		fileSize: fileSizeEmptyValue,
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		PrepareFunction: cmd.Prepare,
		CheckFunction:   cmd.Check,
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "file-size",
		Short: "get sourceFile size from Byte",
		Long:  "",
	})

	root.AddCMD(cmd)
	return cmd
}
