package credential_cmd

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/param_cmd"
	"qshell/iqshell/credential"
	"qshell/qn_error"
)

type removeCMD struct {
	*param_cmd.ParamCMD

	name string
}

func (cmd *removeCMD) Check(context *common.QShellContext) qn_error.IError {
	if cmd.name == "" {
		return qn_error.NewInvalidUserParamError("name can not empty")
	}
	return nil
}

func (cmd *removeCMD) Execute(context *common.QShellContext) qn_error.IError {
	return credential.RemoveCredential(cmd.name)
}

func loadUserRemoveCMD(root param_cmd.IParamCMD) {
	cmd := &removeCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "remove",
		Short: "remove credential",
		Long:  "",
	})

	cmd.FlagsStringVar(&cmd.name, "name", "", "", "credential name")

	root.AddCMD(cmd)
}
