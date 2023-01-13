package env

import (
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/contract"
)

type HadeEnvProvider struct {
	BaseFolder string
}

func (hep *HadeEnvProvider) Register(container framework.Container) framework.NewInstance {
	return NewHadeEnvService
}

func (hep *HadeEnvProvider) Name() string {
	return contract.EnvKey
}

func (hep *HadeEnvProvider) IsDefer() bool {
	return false
}

func (hep *HadeEnvProvider) Boot(container framework.Container) error {
	return nil
}

func (hep *HadeEnvProvider) Params(container framework.Container) []interface{} {
	return []interface{}{"./app.env"}
}
