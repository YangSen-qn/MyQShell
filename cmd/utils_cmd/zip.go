package utils_cmd

import (
	"os"
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/param_cmd"
	"qshell/qn_error"
	"qshell/qn_util"
)

type unzipCMD struct {
	*param_cmd.ParamCMD

	sourceFile     string
	destinationDir string
}

func (cmd *unzipCMD) Prepare(context *common.QShellContext) qn_error.IError {
	cmd.sourceFile = cmd.GetFirstArg()
	if cmd.destinationDir == "" {
		destinationDir, err := os.Getwd()
		if err != nil {
			return qn_error.NewFilePathError(err.Error())
		} else {
			cmd.destinationDir = destinationDir
		}
	}
	return nil
}

func (cmd *unzipCMD) Check(context *common.QShellContext) qn_error.IError {
	if cmd.sourceFile == "" {
		return qn_error.NewInvalidUserParamError("sourceFile path can not empty")
	} else {
		return nil
	}
}

func (cmd *unzipCMD) Execute(context *common.QShellContext) qn_error.IError {
	return qn_util.Unzip(cmd.sourceFile, cmd.destinationDir)
}

func loadZipCMD(root param_cmd.IParamCMD) param_cmd.IParamCMD {
	cmd := &unzipCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		PrepareFunction: cmd.Prepare,
		CheckFunction:   cmd.Check,
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "unzip",
		Short: "unzip file",
		Long:  "",
	})

	cmd.FlagsStringVar(&cmd.destinationDir, "out", "", "", "destination dir by unzip")

	root.AddCMD(cmd)
	return cmd
}