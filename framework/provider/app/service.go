package app

import (
	"errors"
	"flag"
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/util"
	"path/filepath"
)

type HadeApp struct {
	container  framework.Container
	baseFolder string
	configMap  map[string]string
}

// Version 实现版本
func (ha HadeApp) Version() string {
	return "0.0.3"
}

func (ha HadeApp) BaseFolder() string {
	if ha.baseFolder != "" {
		return ha.baseFolder
	}
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数，默认是当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}
	return util.GetExecDirectory()
}

// ConfigFolder  表示配置文件地址
func (h HadeApp) ConfigFolder() string {
	if value, ok := h.configMap["config_folder"]; ok {
		return value
	}
	return filepath.Join(h.BaseFolder(), "config")
}

// LogFolder 表示日志存放地址
func (h HadeApp) LogFolder() string {
	if value, ok := h.configMap["log_folder"]; ok {
		return value
	}
	return filepath.Join(h.StorageFolder(), "log")
}

func (h HadeApp) HttpFolder() string {
	if value, ok := h.configMap["http_folder"]; ok {
		return value
	}
	return filepath.Join(h.BaseFolder(), "http")
}

func (h HadeApp) ConsoleFolder() string {
	if value, ok := h.configMap["console_folder"]; ok {
		return value
	}
	return filepath.Join(h.BaseFolder(), "console")
}

func (h HadeApp) StorageFolder() string {
	if value, ok := h.configMap["storage_folder"]; ok {
		return value
	}
	return filepath.Join(h.BaseFolder(), "storage")
}

// ProviderFolder 定义业务自己的服务提供者地址
func (h HadeApp) ProviderFolder() string {
	if value, ok := h.configMap["provider_folder"]; ok {
		return value
	}
	return filepath.Join(h.BaseFolder(), "provider")
}

// MiddlewareFolder 定义业务自己定义的中间件
func (h HadeApp) MiddlewareFolder() string {
	if value, ok := h.configMap["middleware_folder"]; ok {
		return value
	}
	return filepath.Join(h.HttpFolder(), "middleware")
}

// CommandFolder 定义业务定义的命令
func (h HadeApp) CommandFolder() string {
	if value, ok := h.configMap["command_folder"]; ok {
		return value
	}
	return filepath.Join(h.ConsoleFolder(), "command")
}

// RuntimeFolder 定义业务的运行中间态信息
func (h HadeApp) RuntimeFolder() string {
	if value, ok := h.configMap["runtime_folder"]; ok {
		return value
	}
	return filepath.Join(h.StorageFolder(), "runtime")
}

// TestFolder 定义测试需要的信息
func (h HadeApp) TestFolder() string {
	if value, ok := h.configMap["test_folder"]; ok {
		return value
	}
	return filepath.Join(h.BaseFolder(), "test")
}

func NewHadeAppService(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &HadeApp{container: container, baseFolder: baseFolder}, nil
}

func (ha *HadeApp) LoadAppConfig(kv map[string]string) {
	for key, val := range kv {
		ha.configMap[key] = val
	}
}
