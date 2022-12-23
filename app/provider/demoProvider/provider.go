package demoProvider

import (
	"fmt"
	"github.com/godemo/coredemo/framework"
)

type DemoProvider struct {
	framework.ServiceProvider
}

func (demoProvider *DemoProvider) Name() string {
	return DemoBizKey
}

func (demoProvider *DemoProvider) Register(c framework.Container) framework.NewInstance {
	return NewDemoService
}

func (demoProvider *DemoProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

func (demoProvider *DemoProvider) IsDefer() bool {
	return true
}

func (demoProvider *DemoProvider) Boot(c framework.Container) error {
	// log
	fmt.Println("prepare to bind " + DemoBizKey)
	return nil
}
