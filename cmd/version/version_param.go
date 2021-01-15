package version

import (
	"../../execute"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Run:   runFunction,
}

func AddCommand(superCommand *cobra.Command) {
	superCommand.AddCommand(versionCmd)
}

func runFunction(cmd *cobra.Command, params []string) {

	command := &versionCMD{}
	execute.Execute(command)
}
