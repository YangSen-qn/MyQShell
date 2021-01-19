package execute

import (
	"qshell/common"
	"qshell/output"
)

type IExecute interface {
	GetOutput() output.IOutput
	SetOutput(output output.IOutput)

	Check(context *common.QShellContext) common.IQShellError
	Prepare(context *common.QShellContext) common.IQShellError
	Execute(context *common.QShellContext) common.IQShellError
}

func Execute(exe IExecute, context *common.QShellContext) {
	if exe == nil {
		return
	}

	// 配置output
	config := context.GetConfig()
	exe.SetOutput(output.Output(config))
	// 全局配置拉取，比如uc query

	// 检测参数等信息，根据context和用户输入参数，自动匹配command的执行环境，如果不满足执行条件，则返回错误
	if exe.Check(context) != nil {
		return
	}

	// 准备
	if exe.Prepare(context) != nil {
		return
	}

	// 执行
	if exe.Execute(context) != nil {
		return
	}
}
