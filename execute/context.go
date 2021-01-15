package execute

import (
	"../common"
	"../output"
)

type IContext interface {
	output.IOutput

	SetOutput(output output.IOutput)
	GetOutput() output.IOutput

	SetConfig(config *common.Config)
	GetConfig() *common.Config
}

type Context struct {
	output output.IOutput
	config *common.Config
}

func (c *Context) SetOutput(output output.IOutput) {
	c.output = output
}

func (c *Context) GetOutput() output.IOutput {
	return c.output
}

func (c *Context) Output(outputType output.OutputType, data output.IOutputData, err error) {
	if c.output != nil {
		c.output.Output(outputType, data, err)
	}
}

func (c *Context) SetConfig(config *common.Config) {
	c.config = config
}

func (c *Context) GetConfig() *common.Config {
	return c.config
}
