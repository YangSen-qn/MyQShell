package utils_cmd

import (
	"os"
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
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
	filePath := cmd.sourceFile
	qn_util.Unzip(cmd.sourceFile, cmd.destinationDir)
	output_utils.OutputResultWithString(cmd, filePath)
	return nil
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

	root.AddCMD(cmd)
	return cmd
}