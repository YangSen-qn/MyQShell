package credential_cmd

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/param_cmd"
	"qshell/iqshell/credential"
	"qshell/qn_error"
)

type addCMD struct {
	*param_cmd.ParamCMD

	name      string
	accessKey string
	secretKey string
}

func (cmd *addCMD) Check(context *common.QShellContext) qn_error.IError {
	if cmd.name == "" || cmd.accessKey == "" || cmd.secretKey == "" {
		return qn_error.NewInvalidUserParamError("name or accessKey or secretKey invalid")
	}
	return nil
}

func (cmd *addCMD) Execute(context *common.QShellContext) qn_error.IError {
	cred := &credential.Credential{
		Name:      cmd.name,
		AccessKey: cmd.accessKey,
		SecretKey: cmd.secretKey,
	}
	return credential.AddCredential(cred)
}

func loadUserAddCMD(root param_cmd.IParamCMD) {
	cmd := &addCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		CheckFunction:   cmd.Check,
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:     "add",
		Short:   "add credential",
		Long:    "",
		Example: "qshell credential add name \"credential_name\" accessKey \"ak\" secretKey \"sk\"",
	})

	cmd.FlagsStringVar(&cmd.name, "name", "", "", "credential name")
	cmd.FlagsStringVar(&cmd.accessKey, "accessKey", "", "", "access key")
	cmd.FlagsStringVar(&cmd.secretKey, "secretKey", "", "", "secret key")

	root.AddCMD(cmd)
}
