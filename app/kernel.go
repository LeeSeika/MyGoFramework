package http

import (
	"github.com/godemo/coredemo/app/http"
	"github.com/godemo/coredemo/framework/gin"
)

func NewHttpEngine() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	http.Routes(engine)
	return engine, nil
}
