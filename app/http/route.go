package http

import (
	"github.com/godemo/coredemo/app/http/module/demo"
	"github.com/godemo/coredemo/framework/gin"
)

func Routes(engine *gin.Engine) {
	engine.Static("/dist/", "dist")
	demo.Register(engine)
}
