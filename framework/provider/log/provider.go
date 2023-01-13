package log

import (
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/contract"
	"github.com/godemo/coredemo/framework/provider/log/formatter"
	"github.com/godemo/coredemo/framework/provider/log/services"
	"io"
	"strings"
)

type HadeLogProvider struct {
	Driver     string
	Level      contract.LogLevel
	Formatter  contract.Formatter
	CtxFielder contract.CtxFielder
	Output     io.Writer
}

func (hlp *HadeLogProvider) IsDefer() bool {
	return false
}

func (hlp *HadeLogProvider) Name() string {
	return contract.LogKey
}

func (hlp *HadeLogProvider) Boot(container framework.Container) error {
	return nil
}

func (hlp *HadeLogProvider) Register(container framework.Container) framework.NewInstance {
	if hlp.Driver == "" {
		configService := container.MustMake(contract.ConfigKey).(contract.Config)
		logType := strings.ToLower(configService.GetString("log.driver"))
		hlp.Driver = logType
	}
	switch hlp.Driver {
	case "console":
		return services.NewHadeConsoleLog
	case "single":
		return services.NewHadeSingleLog
	case "rotate":
		return services.NewHadeRotateLog
	case "custom":
		return services.NewHadeCustomLog
	}
	return services.NewHadeConsoleLog
}

func logLevel(level string) contract.LogLevel {
	var logLevel contract.LogLevel
	switch strings.ToLower(level) {
	case "panic":
		logLevel = contract.PanicLevel
	case "fatal":
		logLevel = contract.FatalLevel
	case "error":
		logLevel = contract.ErrorLevel
	case "warn":
		logLevel = contract.WarnLevel
	case "info":
		logLevel = contract.InfoLevel
	case "debug":
		logLevel = contract.DebugLevel
	case "trace":
		logLevel = contract.TraceLevel
	default:
		logLevel = contract.InfoLevel
	}
	return logLevel
}

func (hlp *HadeLogProvider) Params(container framework.Container) []interface{} {
	configService := container.MustMake(contract.ConfigKey).(contract.Config)

	if hlp.Level == contract.UnknownLevel {
		hlp.Level = contract.InfoLevel
		if configService.IsExist("log.level") {
			level := configService.GetString("log.level")
			hlp.Level = logLevel(level)
		}
	}

	if hlp.Formatter == nil {
		// default
		hlp.Formatter = formatter.TextFormatter
		if configService.IsExist("log.formatter") {
			fmtStr := configService.GetString("log.formatter")
			if fmtStr == "text" {
				hlp.Formatter = formatter.TextFormatter
			} else if fmtStr == "json" {
				// TODO

			}
		}
	}

	return []interface{}{container, hlp.Level, hlp.CtxFielder, hlp.Formatter}
}
