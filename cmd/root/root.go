package root

import (
	"qshell/cmd/credential"
	"qshell/cmd/param"
	"qshell/cmd/path"
	"qshell/cmd/utils"
	"qshell/cmd/version"

	"qshell/cmd/common"
)

var (
	config = &common.Config{}
)

type RootCMD struct {
	*param.ParamCMD
}

var rootCmd *RootCMD

func init() {
	rootCmd = &RootCMD{
		ParamCMD: param.NewParamCMD(),
	}

	rootCmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:                    "qshell",
		Short:                  "",
		Long:                   "",
		Version:                common.GetVersion(),
		BashCompletionFunction: "",
	})

	configSubCMD()
}

func configSubCMD() {
	version.LoadCMD(rootCmd)
	credential.LoadCMD(rootCmd)
	path.LoadCMD(rootCmd)
	utils.LoadCMD(rootCmd)
}

func Execute() error {
	return rootCmd.CobraExecute()
}
