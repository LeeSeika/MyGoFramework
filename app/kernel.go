package http

import (
	"github.com/godemo/coredemo/app/http"
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/gin"
)

func NewHttpEngine(container framework.Container) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.SetContainer(container)
	http.Routes(engine)
	return engine, nil
}
