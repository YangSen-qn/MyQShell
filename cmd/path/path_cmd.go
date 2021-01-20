package path

import (
	"qshell/cmd/param"
)

type pathCMD struct {
	*param.ParamCMD
}

func configThePathCMD(root param.IParamCMD) param.IParamCMD  {
	cmd := &pathCMD{
		ParamCMD: param.NewParamCMD(),
	}

	//cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
	//	ExecuteFunction: cmd.Execute,
	//})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:                    "path",
		Short:                  "manager path",
		Long:                   "",
	})

	root.AddCMD(cmd)
	return cmd
}

func ConfigCMD(root param.IParamCMD) {
	path := configThePathCMD(root)
	configRootPathCMD(path)
	configConfigPathCMD(path)
}