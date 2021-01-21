package path_cmd

import (
	"qshell/cmd/param_cmd"
)

type pathCMD struct {
	*param_cmd.ParamCMD
}

func configThePathCMD(root param_cmd.IParamCMD) param_cmd.IParamCMD  {
	cmd := &pathCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	//cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
	//	ExecuteFunction: cmd.Execute,
	//})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:                    "path",
		Short:                  "manager path",
		Long:                   "",
	})

	root.AddCMD(cmd)
	return cmd
}

func LoadCMD(root param_cmd.IParamCMD) {
	path := configThePathCMD(root)
	loadRootPathCMD(path)
	loadConfigPathCMD(path)
}