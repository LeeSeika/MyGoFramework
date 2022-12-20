package demoProvider

import (
	"fmt"
	"github.com/godemo/coredemo/framework"
)

// demo服务的实现
type DemoService struct {
	Service
	c framework.Container
}

func NewDemoService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	fmt.Println("new demo service")
	return &DemoService{c: container}, nil
}
