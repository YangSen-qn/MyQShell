package utils

import (
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param"
	"qshell/qn_error"
	"qshell/util"
)

type etagCMD struct {
	*param.ParamCMD

	filePath string
}

func (cmd *etagCMD) Prepare(context *common.QShellContext) error {
	cmd.filePath = cmd.GetFirstArg()
	return nil
}

func (cmd *etagCMD) Check(context *common.QShellContext) error {
	if cmd.filePath == "" {
		return qn_error.NewInvalidUserParamError("sourceFile path can not empty")
	} else {
		return nil
	}
}

func (cmd *etagCMD) Execute(context *common.QShellContext) error {
	etag, err := util.GetEtag(cmd.filePath)
	if err != nil {
		return err
	}

	output_utils.OutputResultWithString(cmd, etag)
	return nil
}

func loadEtagCMD(root param.IParamCMD) param.IParamCMD {
	cmd := &etagCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		PrepareFunction: cmd.Prepare,
		CheckFunction:   cmd.Check,
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "etag <LocalFilePath>",
		Short: "Calculate the hash of local sourceFile using the algorithm of qiniu qetag",
		Long:  "",
	})

	root.AddCMD(cmd)
	return cmd
}
