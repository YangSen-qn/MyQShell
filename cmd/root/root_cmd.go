package root

import (
	"github.com/spf13/cobra"
	"qshell/common"
	"qshell/execute"

	_ "qshell/cmd/user"
	_ "qshell/cmd/version"
)

var (
	config = &common.Config{}
)

var RootCmd = &cobra.Command{
	Use:     "qshell",
	Short:   "Qiniu commandline tool for managing your bucket and CDN",
	Version: common.GetVersion(),
	Run:     runFunction,
}

func init() {

	RootCmd.Flags().StringVarP(&config.OutputFormatValue, "outputFormat", "", "", "")
}

func runFunction(cmd *cobra.Command, args []string) {

	context := execute.NewQShellContext(cmd.Context())
	context.SetConfig(config)
	cmd.Context()
	RootCmd.Flags().StringVarP(&config.OutputFormatValue, "outputFormat", "", "", "")
}
