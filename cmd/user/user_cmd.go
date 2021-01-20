package user

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/param"
	"qshell/iqshell/user"
	"qshell/qn_shell_error"
)

type userCMD struct {
	*param.ParamCMD

	name string
}

func (cmd *userCMD) Execute(context *common.QShellContext) qn_shell_error.IQShellError {
	if cmd.name == "" {
		currentCredential := user.CurrentCredential()
		if currentCredential == nil {
			return qn_shell_error.NewInvalidWarningError("current credential not exist")
		}
		output.OutputResult(cmd, currentCredential)
	} else {
		credential := user.GetCredential(cmd.name)
		if credential == nil {
			return qn_shell_error.NewInvalidWarningError("credential not exist for name:" + cmd.name)
		}
		output.OutputResult(cmd, credential)
	}

	return nil
}

func configUserCMD(root param.IParamCMD) param.IParamCMD {
	cmd := &userCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "credential",
		Short: "manager credential",
		Long:  "",
	})

	cmd.FlagsStringVar(&cmd.name, "name", "", "", "credential name")

	root.AddCMD(cmd)
	return cmd
}

func ConfigCMD(root param.IParamCMD) {
	userCMD := configUserCMD(root)
	configUserListCMD(userCMD)
	configUserAddCMD(userCMD)
	configUserRemoveCMD(userCMD)
	configUserChangeCMD(userCMD)
}
