package utils_cmd

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param_cmd"
	"qshell/qn_error"
	"qshell/qn_util"
)

type base64CMD struct {
	*param_cmd.ParamCMD

	value     string
	isUrlSafe bool
	isEncode  bool
	isDecode  bool
}

func (cmd *base64CMD) Prepare(context *common.QShellContext) qn_error.IError {
	cmd.value = cmd.GetFirstArg()
	return nil
}

func (cmd *base64CMD) Check(context *common.QShellContext) qn_error.IError {
	if cmd.value == "" {
		return qn_error.NewInvalidUserParamError("base64 data can not be empty")
	} else {
		return nil
	}
}

func (cmd *base64CMD) Execute(context *common.QShellContext) qn_error.IError {
	if cmd.isEncode {
		result, err := qn_util.Base64Encode(cmd.value, cmd.isUrlSafe)
		if err != nil {
			return err
		} else {
			output_utils.OutputResultWithString(cmd, result)
		}
	} else if cmd.isDecode {
		result := qn_util.Base64Decode(cmd.value, cmd.isUrlSafe)
		output_utils.OutputResultWithString(cmd, result)
	}
	return nil
}

func loadBase64CMD(root param_cmd.IParamCMD) param_cmd.IParamCMD {
	cmd := &base64CMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "base64",
		Short: "Base64 string, default nor value safe",
		Long:  "",
	})

	cmd.FlagsBoolVar(&cmd.isUrlSafe, "url-safe", "", true, "url safe base64 encode or decode, default false")
	cmd.FlagsBoolVar(&cmd.isEncode, "encode", "", true, "base64 encode, default true")
	cmd.FlagsBoolVar(&cmd.isDecode, "decode", "", false, "base64 decode, default false")

	root.AddCMD(cmd)
	return cmd
}
