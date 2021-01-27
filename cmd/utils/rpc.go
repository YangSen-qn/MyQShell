package utils

import (
	"bufio"
	"os"
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param"
	"qshell/util"
)

type rpcCMD struct {
	*param.ParamCMD

	url      string
	isEncode bool
	isDecode bool
}

func (cmd *rpcCMD) Execute(context *common.QShellContext) error {
	args := cmd.GetArgs()
	if cmd.isEncode {
		if len(args) > 0 {
			for _, param := range args {
				encodedStr := util.Encode(param)
				output_utils.OutputResultWithString(cmd, encodedStr)
			}
		} else {
			bScanner := bufio.NewScanner(os.Stdin)
			for bScanner.Scan() {
				toEncode := bScanner.Text()
				encodedStr, _ := util.Decode(string(toEncode))
				output_utils.OutputResultWithString(cmd, encodedStr)
			}
		}
	} else if cmd.isDecode {
		if len(args) > 0 {
			for _, param := range args {
				decodedStr, _ := util.Decode(param)
				output_utils.OutputResultWithString(cmd, decodedStr)
			}
		} else {
			bScanner := bufio.NewScanner(os.Stdin)
			for bScanner.Scan() {
				toDecode := bScanner.Text()
				decodedStr, _ := util.Decode(string(toDecode))
				output_utils.OutputResultWithString(cmd, decodedStr)
			}
		}
	}
	return nil
}

func loadRPCCMD(root param.IParamCMD) param.IParamCMD {
	cmd := &rpcCMD{
		ParamCMD: param.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param.ParamCMDConfig{
		Use:   "rpc",
		Short: "rpc encode and decode of qiniu",
		Long:  "",
	})

	cmd.FlagsBoolVar(&cmd.isEncode, "encode", "", true, "rpc encode")
	cmd.FlagsBoolVar(&cmd.isDecode, "decode", "", false, "rpc decode")

	root.AddCMD(cmd)
	return cmd
}
