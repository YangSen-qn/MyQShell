package credential

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/param"
	"qshell/iqshell/credential"
	"qshell/qn_error"
)

type changeCMD struct {
	*param.ParamCMD

	name string
}

func (cmd *changeCMD) Check(context *common.QShellContext) error {
	if cmd.name == "" {
		return qn_error.NewInvalidUserParamError("name can not empty")
	}
	return nil
}

func (cmd *changeCMD) Execute(context *common.QShellContext) error {
	return credential.SetCurrentCredential(cmd.name)
}

func loadUserChangeCMD(root param.IParamCMD) {
	cmd := &changeCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		CheckFunction:   cmd.Check,
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "change",
		Short: "change current credential",
		Long:  "",
	})

	cmd.FlagsStringVar(&cmd.name, "name", "", "", "credential name")

	root.AddCMD(cmd)
}
