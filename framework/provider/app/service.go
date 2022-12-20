package app

import (
	"flag"
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/util"
)

type HadeApp struct {
	container  framework.Container
	baseFolder string
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
