package http

import (
	"github.com/godemo/coredemo/app/http/module/demo"
	"github.com/godemo/coredemo/framework/gin"
	"github.com/godemo/coredemo/framework/middleware/static"
)

func Routes(engine *gin.Engine) {
	engine.Use(static.Serve("/", static.LocalFile("./dist", false)))
	demo.Register(engine)
}
