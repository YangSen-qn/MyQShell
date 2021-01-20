package execute

import (
	"qshell/cmd/common"
	"qshell/cmd/output"
	error2 "qshell/qn_shell_error"
)

type CommandFunction func(context *common.QShellContext) error2.IQShellError

type CommandConfig struct {
	Output          output.IOutput
	CheckFunction   CommandFunction
	PrepareFunction CommandFunction
	ExecuteFunction CommandFunction
}

type Command struct {
	output output.IOutput

	checkFunction   CommandFunction
	prepareFunction CommandFunction
	executeFunction CommandFunction
}

func NewCommand() *Command {
	return &Command{
		output:          nil,
		checkFunction:   nil,
		prepareFunction: nil,
		executeFunction: nil,
	}
}
func (c *Command) ConfigCommand(param CommandConfig) {
	c.output = param.Output
	c.checkFunction = param.CheckFunction
	c.prepareFunction = param.PrepareFunction
	c.executeFunction = param.ExecuteFunction
}

func (c *Command) SetOutput(output output.IOutput) {
	c.output = output
}

func (c *Command) GetOutput() output.IOutput {
	return c.output
}

func (c *Command) Check(context *common.QShellContext) error2.IQShellError {
	if c.checkFunction == nil {
		return nil
	} else {
		return c.checkFunction(context)
	}
}

func (c *Command) Prepare(context *common.QShellContext) error2.IQShellError {
	if c.prepareFunction == nil {
		return nil
	} else {
		return c.prepareFunction(context)
	}
}

func (c *Command) Execute(context *common.QShellContext) error2.IQShellError {
	if c.executeFunction == nil {
		return nil
	} else {
		return c.executeFunction(context)
	}
}

func (c *Command) Output(outputType output.OutputType, data output.IOutputData, err error2.IQShellError) {
	if c.output != nil {
		c.output.Output(outputType, data, err)
	}
}
