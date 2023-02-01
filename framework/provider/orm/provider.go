package orm

import (
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/contract"
)

type HadeOrmProvider struct {
}

func (hop *HadeOrmProvider) Register() framework.NewInstance {
	return NewHadeOrmService
}

func (hop *HadeOrmProvider) Boot(container framework.Container) error {
	return nil
}

func (hop *HadeOrmProvider) Params(container framework.Container) []interface{} {
	return []interface{}{container}
}

func (hop *HadeOrmProvider) IsDefer() bool {
	return true
}

func (hop *HadeOrmProvider) Name() string {
	return contract.ORMKey
}
