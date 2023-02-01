package orm

import (
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/contract"
)

// WithDryRun 设置空跑模式
func WithDryRun() contract.DBOption {
	return func(container framework.Container, config *contract.DBConfig) error {
		config.DryRun = true
		return nil
	}
}

func WithConfigPath(configPath string) contract.DBOption {
	return func(container framework.Container, config *contract.DBConfig) error {
		configService := container.MustMake(contract.ConfigKey).(contract.Config)
		err := configService.Load(configPath, config)
		if err != nil {
			return err
		}
		return nil
	}
}
