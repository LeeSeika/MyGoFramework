package app

import (
	"errors"
	"github.com/godemo/coredemo/framework"
)

type HadeAppProvider struct {
	BaseFolder string
}

func (hap *HadeAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, hap.BaseFolder}
}

func NewHadeApp(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &HadeApp{container: container, baseFolder: baseFolder}, nil
}
