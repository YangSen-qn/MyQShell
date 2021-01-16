package execute

import (
	"qshell/output"
)

type ICheck interface {
	Check() error
}

type IPrepare interface {
	Prepare() error
}

type IExecute interface {
	Execute() error
}

func Execute(cmd ICommand) {
	if cmd == nil {
		return
	}

	config := cmd.GetConfig()

	cmd.SetOutput(output.Output(config))

	check, isOk := cmd.(ICheck)
	if isOk {
		if err := check.Check(); err != nil {
			return
		}
	}

	prepare, isOk := cmd.(IPrepare)
	if isOk {
		if err := prepare.Prepare(); err != nil {
			return
		}
	}

	execute, isOk := cmd.(IExecute)
	if isOk {
		_ = execute.Execute()
	}
}
