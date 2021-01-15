package root

import (
	"../../common"
	"../user"
	"../version"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "qshell",
	Short:   "Qiniu commandline tool for managing your bucket and CDN",
	Version: common.GetVersion(),
}

func init() {

	version.AddCommand(RootCmd)
	user.AddCommand(RootCmd)
}
