package utils_cmd

import (
	"fmt"
	"qshell/cmd/common"
	"qshell/cmd/execute"
	"qshell/cmd/output/output_utils"
	"qshell/cmd/param_cmd"
	"qshell/qn_error"
	"qshell/qn_util"
	"strconv"
)

type dateCMD struct {
	*param_cmd.ParamCMD

	isSecond  bool
	isMilli   bool
	isNano    bool
	before    int64
	timestamp int64
}

func (cmd *dateCMD) Execute(context *common.QShellContext) qn_error.IError {

	var before int64
	if cmd.isNano {
	} else if cmd.isMilli {
		before = cmd.before * 1e6
	} else {
		before = cmd.before * 1e9
	}

	var timestamp int64
	if cmd.timestamp == -1 {
		timestamp = qn_util.Timestamp()
		fmt.Println("=== A")
	} else {
		timestamp = cmd.timestamp
		fmt.Println("=== B:", strconv.Itoa(int(cmd.timestamp)))
	}

	date := qn_util.DateWithTimestamp(timestamp, -1*before)
	output_utils.OutputResultWithString(cmd, date)
	return nil
}

func loadDateCMD(root param_cmd.IParamCMD) param_cmd.IParamCMD {
	cmd := &dateCMD{
		ParamCMD: param_cmd.NewParamCMD(),
	}

	cmd.ConfigParamCMDExecuteConfig(execute.CommandConfig{
		ExecuteFunction: cmd.Execute,
	})

	cmd.ConfigParamCMDParseConfig(param_cmd.ParamCMDConfig{
		Use:   "date",
		Short: "show date by seconds or milli seconds or nano seconds",
		Long:  "",
	})

	cmd.FlagsInt64Var(&cmd.timestamp, "timestamp", "", -1, "get date by timestamp")
	cmd.FlagsInt64Var(&cmd.before, "before", "", 0, "value of date before, get date by timestamp is set or set by current")
	cmd.FlagsBoolVar(&cmd.isSecond, "seconds", "", true, "unit is seconds")
	cmd.FlagsBoolVar(&cmd.isMilli, "milli", "", false, "unit is milli second")
	cmd.FlagsBoolVar(&cmd.isNano, "nano", "", false, "unit is nano seconds")

	root.AddCMD(cmd)
	return cmd
}
