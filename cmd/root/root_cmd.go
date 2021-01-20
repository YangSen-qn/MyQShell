package root

import (
	"qshell/cmd/param"
	"qshell/cmd/path"
	"qshell/cmd/user"
	"qshell/cmd/version"

	"qshell/cmd/common"
)

var (
	config = &common.Config{}
)

type RootCMD struct {
	*param.ParamCMD

	config *common.Config
}

var rootCmd *RootCMD

func init() {
	rootCmd = &RootCMD{
		ParamCMD: param.NewParamCMD(),
		config: &common.Config{},
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
	version.ConfigCMD(rootCmd)
	user.ConfigCMD(rootCmd)
	path.ConfigCMD(rootCmd)
}

func Execute() error {
	return rootCmd.CobraExecute()
}