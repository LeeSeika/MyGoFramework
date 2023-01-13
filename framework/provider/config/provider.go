package config

import (
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/contract"
	"path/filepath"
)

type HadeConfigProvider struct {
}

func (hcp *HadeConfigProvider) Params(c framework.Container) []interface{} {
	envService := c.MustMake(contract.EnvKey).(contract.Env)
	appService := c.MustMake(contract.AppKey).(contract.App)
	configFolder := appService.ConfigFolder()
	appEnv := envService.AppEnv()
	envFolder := filepath.Join(configFolder, appEnv)
	return []interface{}{c, envFolder, envService.All()}
}

func (hcp *HadeConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeConfig
}

// Boot will called when the service instantiate
func (hcp *HadeConfigProvider) Boot(c framework.Container) error {
	return nil
}

// IsDefer define whether the service instantiate when first make or register
func (hcp *HadeConfigProvider) IsDefer() bool {
	return false
}

func (hcp *HadeConfigProvider) Name() string {
	return contract.ConfigKey
}
