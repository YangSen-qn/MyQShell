package common

import (
	"context"
	"time"
)

type QShellContext struct {
	context context.Context

	config *Config
}

func NewQShellContext(context context.Context) *QShellContext {
	return &QShellContext{
		context: context,
	}
}

func (c *QShellContext) SetConfig(config *Config) {
	c.config = config
}

func (c *QShellContext) GetConfig() *Config {
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
