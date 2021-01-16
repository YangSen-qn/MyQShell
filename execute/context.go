package execute

import (
	"context"
	"qshell/common"
	"time"
)

type IQShellContext interface {
	context.Context

	SetConfig(config *common.Config)
	GetConfig() *common.Config
}

type QShellContext struct {
	context context.Context

	config *common.Config
}

func NewQShellContext(context context.Context) *QShellContext {
	return &QShellContext{
		context: context,
	}
}

func (c *QShellContext) SetConfig(config *common.Config) {
	c.config = config
}

func (c *QShellContext) GetConfig() *common.Config {
	return c.config
}


func (c *QShellContext) Deadline() (deadline time.Time, ok bool) {
	return c.context.Deadline()
}

func (c *QShellContext) Done() <-chan struct{} {
	return c.context.Done()
}

func (c *QShellContext) Err() error {
	return c.context.Err()
}

func (c *QShellContext) Value(key interface{}) interface{} {
	return c.context.Value(key)
}
