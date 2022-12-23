package gin

import (
	"github.com/godemo/coredemo/framework"
	"golang.org/x/net/context"
)

func (c *Context) BaseContext() context.Context {
	return c.Request.Context()
}

func (c *Context) Make(key string) (interface{}, error) {
	return c.container.Make(key)
}

func (c *Context) MustMake(key string) interface{} {
	return c.container.MustMake(key)
}

func (c *Context) MakeNew(key string, params []interface{}) (interface{}, error) {
	return c.container.MakeNew(key, params)
}

// engine

func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}

func (engine *Engine) SetContainer(container framework.Container) {
	engine.container = container
}
