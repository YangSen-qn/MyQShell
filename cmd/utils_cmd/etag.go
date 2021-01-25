package utils_cmd

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/output/message"
	"qshell/cmd/param_cmd"
	"qshell/qn_error"
	"qshell/qn_util"
)

type etagCMD struct {
	*param_cmd.ParamCMD

	filePath string
}

func (cmd *etagCMD) Prepare(context *common.QShellContext) qn_error.IError {
	cmd.filePath = cmd.GetFirstArg()
	return nil
}

func (cmd *etagCMD) Check(context *common.QShellContext) qn_error.IError {
	if cmd.filePath == "" {
		return qn_error.NewInvalidUserParamError("sourceFile path can not empty")
	} else {
		return nil
	}
}

func (cmd *etagCMD) Execute(context *common.QShellContext) qn_error.IError {
	etag, err := qn_util.GetEtag(cmd.filePath)
	if err != nil {
		return qn_error.NewIOError(err.Error())
	}
	output.OutputResult(cmd, message.NewStringOutputData(etag))
	return nil
}

func loadEtagCMD(root param_cmd.IParamCMD) param_cmd.IParamCMD {
	cmd := &etagCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		PrepareFunction: cmd.Prepare,
		CheckFunction:   cmd.Check,
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "etag <LocalFilePath>",
		Short: "Calculate the hash of local sourceFile using the algorithm of qiniu qetag",
		Long:  "",
	})

	root.AddCMD(cmd)
	return cmd
}
