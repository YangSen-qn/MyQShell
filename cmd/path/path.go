package path

import (
	"qshell/cmd/param"
)

type pathCMD struct {
	*param.ParamCMD
}

func configThePathCMD(root param.IParamCMD) param.IParamCMD {
	cmd := &pathCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "path",
		Short: "manager path",
		Long:  "",
	})

	root.AddCMD(cmd)
	return cmd
}

func LoadCMD(root param.IParamCMD) {
	path := configThePathCMD(root)
	loadRootPathCMD(path)
	loadConfigPathCMD(path)
}
