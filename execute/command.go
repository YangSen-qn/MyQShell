package execute

import (
	"qshell/output"
)

type ICommand interface {
	output.IOutput

	SetOutput(output output.IOutput)
	GetOutput() output.IOutput
}


type Command struct {
	output output.IOutput
}


func (c *Command) SetOutput(output output.IOutput) {
	c.output = output
}

func (c *Command) GetOutput() output.IOutput {
	return c.output
}

func (c *Command) Output(outputType output.OutputType, data output.IOutputData, err error) {
	if c.output != nil {
		c.output.Output(outputType, data, err)
	}
}