package utils_cmd

import (
	"bufio"
	"os"
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output"
	"qshell/cmd/output/message"
	"qshell/cmd/param_cmd"
	"qshell/qn_error"
	"qshell/qn_util"
)

type rpcCMD struct {
	*param_cmd.ParamCMD

	url      string
	isEncode bool
	isDecode bool
}

func (cmd *rpcCMD) Execute(context *common.QShellContext) qn_error.IError {
	args := cmd.GetArgs()
	if cmd.isEncode {
		if len(args) > 0 {
			for _, param := range args {
				encodedStr := qn_util.Encode(param)
				output.OutputResult(cmd, message.NewStringOutputData(encodedStr))
			}
		} else {
			bScanner := bufio.NewScanner(os.Stdin)
			for bScanner.Scan() {
				toEncode := bScanner.Text()
				encodedStr, _ := qn_util.Decode(string(toEncode))
				output.OutputResult(cmd, message.NewStringOutputData(encodedStr))
			}
		}
	} else if cmd.isDecode {
		if len(args) > 0 {
			for _, param := range args {
				decodedStr, _ := qn_util.Decode(param)
				output.OutputResult(cmd, message.NewStringOutputData(decodedStr))
			}
		} else {
			bScanner := bufio.NewScanner(os.Stdin)
			for bScanner.Scan() {
				toDecode := bScanner.Text()
				decodedStr, _ := qn_util.Decode(string(toDecode))
				output.OutputResult(cmd, message.NewStringOutputData(decodedStr))
			}
		}
	}
	return nil
}

func loadRPCCMD(root param_cmd.IParamCMD) param_cmd.IParamCMD {
	cmd := &rpcCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "rpc",
		Short: "rpc encode and decode of qiniu",
		Long:  "",
	})

	cmd.FlagsBoolVar(&cmd.isEncode, "encode", "", false, "rpc encode")
	cmd.FlagsBoolVar(&cmd.isDecode, "decode", "", false, "rpc decode")

	root.AddCMD(cmd)
	return cmd
}