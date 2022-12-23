package kernel

import (
	"errors"
	"github.com/godemo/coredemo/framework/gin"
	"net/http"
)

type HadeKernelService struct {
	engine *gin.Engine
}

func NewHadeKernelService(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("param error")
	}
	engine := params[0].(*gin.Engine)
	return &HadeKernelService{engine: engine}, nil
}

func (hks *HadeKernelService) HttpEngine() http.Handler {
	return hks.engine
}
