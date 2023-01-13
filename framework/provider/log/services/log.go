package services

import (
	"context"
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/contract"
	"github.com/godemo/coredemo/framework/provider/log/formatter"
	"io"
	"log"
	"time"
)

type HadeLog struct {
	container  framework.Container
	level      contract.LogLevel
	formatter  contract.Formatter
	output     io.Writer
	ctxFielder contract.CtxFielder
}

func (hl *HadeLog) IsLevelEnable(level contract.LogLevel) bool {
	return level == contract.DebugLevel || level == contract.FatalLevel || level == contract.InfoLevel || level == contract.TraceLevel || level == contract.UnknownLevel || level == contract.WarnLevel || level == contract.ErrorLevel
}

func (hl *HadeLog) Info(ctx context.Context, msg string, fields map[string]interface{}) {

}

func (hl *HadeLog) Panic(ctx context.Context, msg string, fields map[string]interface{}) {

}
func (hl *HadeLog) Fatal(ctx context.Context, msg string, fields map[string]interface{}) {

}
func (hl *HadeLog) Error(ctx context.Context, msg string, fields map[string]interface{}) {

}
func (hl *HadeLog) Warn(ctx context.Context, msg string, fields map[string]interface{}) {

}
func (hl *HadeLog) Debug(ctx context.Context, msg string, fields map[string]interface{}) {

}
func (hl *HadeLog) Trace(ctx context.Context, msg string, fields map[string]interface{}) {

}
func (hl *HadeLog) SetLevel(level contract.LogLevel) {

}
func (hl *HadeLog) SetCtxFielder(handler contract.CtxFielder) {

}
func (hl *HadeLog) SetFormatter(formatter contract.Formatter) {

}

func (hl *HadeLog) SetOutput(output io.Writer) {
	hl.output = output
}

func (hl *HadeLog) logf(level contract.LogLevel, ctx context.Context, msg string, fields map[string]interface{}) error {
	if !hl.IsLevelEnable(level) {
		return nil
	}
	fs := fields
	if hl.ctxFielder != nil {
		fielder := hl.ctxFielder(ctx)
		for k, v := range fielder {
			fs[k] = v
		}
	}

	if hl.formatter == nil {
		hl.formatter = formatter.TextFormatter
	}

	bytes, err := hl.formatter(level, time.Now(), msg, fs)
	if err != nil {
		return err
	}

	if hl.level == contract.PanicLevel {
		log.Panicln(string(bytes))
		return nil
	}

	_, err = hl.output.Write(bytes)
	if err != nil {
		return err
	}
	_, err = hl.output.Write([]byte("/r/n"))
	if err != nil {
		return err
	}
	return nil
}
