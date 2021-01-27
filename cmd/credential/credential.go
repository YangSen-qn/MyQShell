package credential

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/param"
	"qshell/iqshell/credential"
	"qshell/qn_error"
)

type credentialCMD struct {
	*param.ParamCMD

	name string
}

func (cmd *credentialCMD) Execute(context *common.QShellContext) error {
	if cmd.name == "" {
		currentCredential := credential.CurrentCredential()
		if currentCredential == nil {
			return qn_error.NewExecuteError("current credential not exist")
		}
		output.OutputResult(cmd, currentCredential)
	} else {
		credential := credential.GetCredential(cmd.name)
		if credential == nil {
			return qn_error.NewExecuteError("credential not exist for name: %s", cmd.name)
		}
		output.OutputResult(cmd, credential)
	}

	return nil
}

func configUserCMD(root param.IParamCMD) param.IParamCMD {
	cmd := &credentialCMD{
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

func LoadCMD(root param.IParamCMD) {
	userCMD := configUserCMD(root)
	loadUserListCMD(userCMD)
	loadUserAddCMD(userCMD)
	loadUserRemoveCMD(userCMD)
	loadUserChangeCMD(userCMD)
}
