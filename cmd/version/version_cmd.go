package version

import (
	"../../common"
	"../../execute"
	"../../output"
)

type versionCMD struct {
	execute.Command
}

func (cmd *versionCMD) Execute() error {
	output.OutputResult(cmd, output.NewStringOutputData(common.GetVersion()))
	return nil
}
