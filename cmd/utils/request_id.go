package utils

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param"
	"qshell/qn_error"
	"qshell/util"
)

type parseRequestIdCMD struct {
	*param.ParamCMD

	requestId string
}

func (cmd *parseRequestIdCMD) Prepare(context *common.QShellContext) error {
	cmd.requestId = cmd.GetFirstArg()
	return nil
}

func (cmd *parseRequestIdCMD) Check(context *common.QShellContext) error {
	if cmd.requestId == "" {
		return qn_error.NewInvalidUserParamError("request id can not empty")
	} else {
		return nil
	}
}

func (cmd *parseRequestIdCMD) Execute(context *common.QShellContext) error {

	result, err := util.ParseReqId(cmd.requestId)
	if err != nil {
		return err
	}
	output_utils.OutputResultWithString(cmd, result)
	return nil
}

func loadParseRequestIdCMD(root param.IParamCMD) param.IParamCMD {
	cmd := &parseRequestIdCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		PrepareFunction: cmd.Prepare,
		CheckFunction:   cmd.Check,
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "request-id <ReqIdToDecode>",
		Short: "Decode qiniu reqid",
		Long:  "",
	})

	root.AddCMD(cmd)
	return cmd
}
