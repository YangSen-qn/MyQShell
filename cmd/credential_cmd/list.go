package credential_cmd

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param_cmd"
	"qshell/iqshell/credential"
	error2 "qshell/qn_error"
)

type listCMD struct {
	*param_cmd.ParamCMD
}

func (cmd *listCMD) Execute(context *common.QShellContext) error2.IError {
	userList := credential.CredentialList()
	if len(userList) == 0 {
		output_utils.OutputResultWithString(cmd, "credential list is empty")
		return nil
	}

	for _, u := range userList {
		output.OutputResult(cmd, u)
	}
	return nil
}

func loadUserListCMD(root param_cmd.IParamCMD) {
	cmd := &listCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "list",
		Short: "list credential",
		Long:  "",
	})

	root.AddCMD(cmd)
}
