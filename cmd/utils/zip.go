package utils

import (
	"os"
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/param"
	"qshell/qn_error"
	"qshell/util"
)

type unzipCMD struct {
	*param.ParamCMD

	sourceFile     string
	destinationDir string
}

func (cmd *unzipCMD) Prepare(context *common.QShellContext) error {
	cmd.sourceFile = cmd.GetFirstArg()
	if cmd.destinationDir == "" {
		destinationDir, err := os.Getwd()
		if err != nil {
			return qn_error.NewErrorWithError(err)
		} else {
			cmd.destinationDir = destinationDir
		}
	}
	return nil
}

func (cmd *unzipCMD) Check(context *common.QShellContext) error {
	if cmd.sourceFile == "" {
		return qn_error.NewInvalidUserParamError("sourceFile path can not empty")
	} else {
		return nil
	}
}

func (cmd *unzipCMD) Execute(context *common.QShellContext) error {
	return util.Unzip(cmd.sourceFile, cmd.destinationDir)
}

func loadZipCMD(root param.IParamCMD) param.IParamCMD {
	cmd := &unzipCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		PrepareFunction: cmd.Prepare,
		CheckFunction:   cmd.Check,
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "unzip",
		Short: "unzip file",
		Long:  "",
	})

	cmd.FlagsStringVar(&cmd.destinationDir, "out", "", "", "destination dir by unzip")

	root.AddCMD(cmd)
	return cmd
}
