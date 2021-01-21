package credential_cmd

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/param_cmd"
	"qshell/iqshell/credential"
	"qshell/qn_error"
)

type changeCMD struct {
	*param_cmd.ParamCMD

	name string
}

func (cmd *changeCMD) Check(context *common.QShellContext) qn_error.IError {
	if cmd.name == "" {
		return qn_error.NewInvalidUserParamError("name can not empty")
	}
	return nil
}

func (cmd *changeCMD) Execute(context *common.QShellContext) qn_error.IError {
	return credential.SetCurrentCredential(cmd.name)
}

func loadUserChangeCMD(root param_cmd.IParamCMD) {
	cmd := &changeCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "change",
		Short: "change current credential",
		Long:  "",
	})

	cmd.FlagsStringVar(&cmd.name, "name", "", "", "credential name")

	root.AddCMD(cmd)
}
