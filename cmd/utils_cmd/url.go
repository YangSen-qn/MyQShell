package utils_cmd

import (
	"net/url"
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param_cmd"
	"qshell/qn_error"
)

type urlCMD struct {
	*param_cmd.ParamCMD

	url      string
	isEncode bool
	isDecode bool
}

func (cmd *urlCMD) Prepare(context *common.QShellContext) qn_error.IError {
	cmd.url = cmd.GetFirstArg()
	return nil
}

func (cmd *urlCMD) Check(context *common.QShellContext) qn_error.IError {
	if cmd.url == "" {
		return qn_error.NewInvalidUserParamError("value can not empty")
	} else {
		return nil
	}
}

func (cmd *urlCMD) Execute(context *common.QShellContext) qn_error.IError {
	urlString := cmd.url
	if cmd.isEncode {
		urlString = url.QueryEscape(urlString)
	} else if cmd.isDecode {
		urlUnescape, err := url.QueryUnescape(urlString)
		if err != nil {
			return qn_error.NewInvalidUserParamError(err.Error())
		}
		urlString = urlUnescape
	}
	output_utils.OutputResultWithString(cmd, urlString)
	return nil
}

func loadUrlCMD(root param_cmd.IParamCMD) param_cmd.IParamCMD {
	cmd := &urlCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		PrepareFunction: cmd.Prepare,
		CheckFunction:   cmd.Check,
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "value",
		Short: "value encode and decode",
		Long:  "",
	})

	cmd.FlagsBoolVar(&cmd.isEncode, "encode", "", false, "value encode")
	cmd.FlagsBoolVar(&cmd.isDecode, "decode", "", false, "value decode")

	root.AddCMD(cmd)
	return cmd
}
