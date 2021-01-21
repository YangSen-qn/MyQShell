package execute

import (
	"qshell/cmd/common"
	"qshell/cmd/output"
	"qshell/iqshell/load"
	error2 "qshell/qn_error"
)

type IExecute interface {
	GetOutput() output.IOutput
	SetOutput(output output.IOutput)

	Check(context *common.QShellContext) error2.IError
	Prepare(context *common.QShellContext) error2.IError
	Execute(context *common.QShellContext) error2.IError
}

func Execute(exe IExecute, context *common.QShellContext) {
	if exe == nil {
		return
	}

	// 配置output
	config := context.GetConfig()
	op := output.Output(config)
	exe.SetOutput(op)

	// 启动本地配置
	err := load.LoadInterQShell(config.LoadConfig)
	if err != nil {
		output.OutputError(op, err)
		return
	}

	// 全局配置拉取，比如uc query

	// 检测参数等信息，根据context和用户输入参数，自动匹配command的执行环境，如果不满足执行条件，则返回错误
	if err = exe.Check(context); err != nil {
		output.OutputError(op, err)
		return
	}

	// 准备
	if err := exe.Prepare(context); err != nil {
		output.OutputError(op, err)
		return
	}

	// 执行
	if err := exe.Execute(context); err != nil {
		output.OutputError(op, err)
		return
	}
}
