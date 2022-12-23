package app

import (
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/contract"
)

type HadeAppProvider struct {
	BaseFolder string
}

func (hap *HadeAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewHadeAppService
}

func (hap *HadeAppProvider) Name() string {
	return contract.AppKey
}

func (hap *HadeAppProvider) IsDefer() bool {
	return false
}

func (hap *HadeAppProvider) Boot(container framework.Container) error {
	return nil
}

func (hap *HadeAppProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container, hap.BaseFolder}
}
