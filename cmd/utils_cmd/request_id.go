package utils_cmd

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param_cmd"
	"qshell/qn_error"
	"qshell/qn_util"
)

type parseRequestIdCMD struct {
	*param_cmd.ParamCMD

	requestId string
}

func (cmd *parseRequestIdCMD) Prepare(context *common.QShellContext) qn_error.IError {
	cmd.requestId = cmd.GetFirstArg()
	return nil
}

func (cmd *parseRequestIdCMD) Check(context *common.QShellContext) qn_error.IError {
	if cmd.requestId == "" {
		return qn_error.NewInvalidUserParamError("request id can not empty")
	} else {
		return nil
	}
}

func (cmd *parseRequestIdCMD) Execute(context *common.QShellContext) qn_error.IError {

	result, err := qn_util.ParseReqId(cmd.requestId)
	if err != nil {
		return err
	}
	output_utils.OutputResultWithString(cmd, result)
	return nil
}

func loadParseRequestIdCMD(root param_cmd.IParamCMD) param_cmd.IParamCMD {
	cmd := &parseRequestIdCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		PrepareFunction: cmd.Prepare,
		CheckFunction:   cmd.Check,
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "request-id <ReqIdToDecode>",
		Short: "Decode qiniu reqid",
		Long:  "",
	})

	root.AddCMD(cmd)
	return cmd
}