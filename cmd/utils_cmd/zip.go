package utils_cmd

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/output/message"
	"qshell/cmd/param_cmd"
	"qshell/qn_error"
)

type unzipCMD struct {
	*param_cmd.ParamCMD

	sourceFile string
}

func (cmd *unzipCMD) Prepare(context *common.QShellContext) qn_error.IError {
	cmd.sourceFile = cmd.GetFirstArg()
	return nil
}

func (cmd *unzipCMD) Check(context *common.QShellContext) qn_error.IError {
	if cmd.sourceFile == "" {
		return qn_error.NewInvalidUserParamError("sourceFile path can not empty")
	} else {
		return nil
	}
}

func (cmd *unzipCMD) Execute(context *common.QShellContext) qn_error.IError {
	filePath := cmd.sourceFile

	output.OutputResult(cmd, message.NewStringOutputData(filePath))
	return nil
}

func loadZipCMD(root param_cmd.IParamCMD) param_cmd.IParamCMD {
	cmd := &unzipCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		PrepareFunction: cmd.Prepare,
		CheckFunction:   cmd.Check,
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "unzip",
		Short: "unzip file",
		Long:  "",
	})

	root.AddCMD(cmd)
	return cmd
}