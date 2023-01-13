package services

import (
	"errors"
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/contract"
	"os"
)

func NewHadeConsoleLog(params ...interface{}) (interface{}, error) {
	if len(params) != 4 {
		return nil, errors.New("params error")
	}
	container := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	consoleLog := &HadeConsoleLog{}
	consoleLog.container = container
	consoleLog.SetLevel(level)
	consoleLog.SetFormatter(formatter)
	consoleLog.SetCtxFielder(ctxFielder)
	// console的输出位置是os.Stdout
	consoleLog.SetOutput(os.Stdout)

	return consoleLog, nil
}

type HadeConsoleLog struct {
	HadeLog
}
