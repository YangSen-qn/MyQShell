package utils

import (
	"net/url"
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param"
	"qshell/qn_error"
)

type urlCMD struct {
	*param.ParamCMD

	url      string
	isEncode bool
	isDecode bool
}

func (cmd *urlCMD) Prepare(context *common.QShellContext) error {
	cmd.url = cmd.GetFirstArg()
	return nil
}

func (cmd *urlCMD) Check(context *common.QShellContext) error {
	if cmd.url == "" {
		return qn_error.NewInvalidUserParamError("value can not empty")
	} else {
		return nil
	}
}

func (cmd *urlCMD) Execute(context *common.QShellContext) error {
	urlString := cmd.url
	if cmd.isEncode {
		urlString = url.QueryEscape(urlString)
	} else if cmd.isDecode {
		urlUnescape, err := url.QueryUnescape(urlString)
		if err != nil {
			return err
		}
		urlString = urlUnescape
	}
	output_utils.OutputResultWithString(cmd, urlString)
	return nil
}

func loadUrlCMD(root param.IParamCMD) param.IParamCMD {
	cmd := &urlCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		PrepareFunction: cmd.Prepare,
		CheckFunction:   cmd.Check,
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "value",
		Short: "value encode and decode",
		Long:  "",
	})

	cmd.FlagsBoolVar(&cmd.isEncode, "encode", "", false, "value encode")
	cmd.FlagsBoolVar(&cmd.isDecode, "decode", "", false, "value decode")

	root.AddCMD(cmd)
	return cmd
}
