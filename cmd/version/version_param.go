package version

import (
	"github.com/spf13/cobra"

	"qshell/execute"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Run:   runFunction,
}

func runFunction(cmd *cobra.Command, params []string) {

	command := &versionCMD{}
	execute.Execute(command)
}
