package kernel

import (
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/contract"
	"github.com/godemo/coredemo/framework/gin"
)

type HadeKernelProvider struct {
	HttpEngine *gin.Engine
}

func (hkp *HadeKernelProvider) Register(container framework.Container) framework.NewInstance {
	return NewHadeKernelService
}
func (hkp *HadeKernelProvider) Name() string {
	return contract.KernelKey
}
func (hkp *HadeKernelProvider) Params(container framework.Container) []interface{} {
	return []interface{}{hkp.HttpEngine}
}
func (hkp *HadeKernelProvider) IsDefer() bool {
	return false
}
func (hkp *HadeKernelProvider) Boot(container framework.Container) error {
	// 判断有没有注入,有的话用注入的,没的话用默认的
	if hkp.HttpEngine == nil {
		hkp.HttpEngine = gin.Default()
	}
	hkp.HttpEngine.SetContainer(container)
	return nil
}
