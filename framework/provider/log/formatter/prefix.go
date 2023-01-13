package formatter

import "github.com/godemo/coredemo/framework/contract"

func Prefix(level contract.LogLevel) string {
	switch level {
	case contract.PanicLevel:
		return "[Panic]"
	case contract.FatalLevel:
		return "[Fatal]"
	case contract.ErrorLevel:
		return "[Error]"
	case contract.WarnLevel:
		return "[Warn]"
	case contract.DebugLevel:
		return "[Debug]"
	case contract.InfoLevel:
		return "[Info]"
	case contract.TraceLevel:
		return "[Trace]"
	}
	return ""
}
