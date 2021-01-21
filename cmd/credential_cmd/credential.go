package credential_cmd

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/param_cmd"
	"qshell/iqshell/credential"
	"qshell/qn_error"
)

type credentialCMD struct {
	*param_cmd.ParamCMD

	name string
}

func (cmd *credentialCMD) Execute(context *common.QShellContext) qn_error.IError {
	if cmd.name == "" {
		currentCredential := credential.CurrentCredential()
		if currentCredential == nil {
			return qn_error.NewInvalidWarningError("current credential not exist")
		}
		output.OutputResult(cmd, currentCredential)
	} else {
		credential := credential.GetCredential(cmd.name)
		if credential == nil {
			return qn_error.NewInvalidWarningError("credential not exist for name:" + cmd.name)
		}
		output.OutputResult(cmd, credential)
	}

	return nil
}

func configUserCMD(root param_cmd.IParamCMD) param_cmd.IParamCMD {
	cmd := &credentialCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "credential",
		Short: "manager credential",
		Long:  "",
	})

	cmd.FlagsStringVar(&cmd.name, "name", "", "", "credential name")

	root.AddCMD(cmd)
	return cmd
}

func LoadCMD(root param_cmd.IParamCMD) {
	userCMD := configUserCMD(root)
	loadUserListCMD(userCMD)
	loadUserAddCMD(userCMD)
	loadUserRemoveCMD(userCMD)
	loadUserChangeCMD(userCMD)
}
