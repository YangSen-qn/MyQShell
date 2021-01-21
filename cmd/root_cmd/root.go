package root_cmd

import (
	"qshell/cmd/param_cmd"
	"qshell/cmd/path_cmd"
	"qshell/cmd/credential_cmd"
	"qshell/cmd/version_cmd"

	"qshell/cmd/common"
)

var (
	config = &common.Config{}
)

type RootCMD struct {
	*param_cmd.ParamCMD
}

var rootCmd *RootCMD

func init() {
	rootCmd = &RootCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	rootCmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:                    "qshell",
		Short:                  "",
		Long:                   "",
		Version:                common.GetVersion(),
		BashCompletionFunction: "",
	})

	configSubCMD()
}

func configSubCMD() {
	version_cmd.LoadCMD(rootCmd)
	credential_cmd.LoadCMD(rootCmd)
	path_cmd.LoadCMD(rootCmd)
}

func Execute() error {
	return rootCmd.CobraExecute()
}