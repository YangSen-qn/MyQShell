package version

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param"
)

type versionCMD struct {
	*param.ParamCMD
}

func (cmd *versionCMD) Execute(context *common.QShellContext) error {
	output_utils.OutputResultWithString(cmd, common.GetVersion())
	return nil
}

func LoadCMD(root param.IParamCMD) {
	cmd := &versionCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "version",
		Short: "show version",
		Long:  "",
	})

	root.AddCMD(cmd)
}
