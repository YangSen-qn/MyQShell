package credential

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param"
	"qshell/iqshell/credential"
)

type listCMD struct {
	*param.ParamCMD
}

func (cmd *listCMD) Execute(context *common.QShellContext) error {
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

func loadUserListCMD(root param.IParamCMD) {
	cmd := &listCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "list",
		Short: "list credential",
		Long:  "",
	})

	root.AddCMD(cmd)
}
