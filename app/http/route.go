package http

import (
  "github.com/godemo/coredemo/app/http/module/demo"
  "github.com/godemo/coredemo/framework/contract"
  "github.com/godemo/coredemo/framework/gin"
  ginSwagger "github.com/godemo/coredemo/framework/middleware/gin-swagger"
  swaggerFiles "github.com/godemo/coredemo/framework/middleware/gin-swagger/swaggerFiles"
  "github.com/godemo/coredemo/framework/middleware/static"
)

func Routes(engine *gin.Engine) {
  engine.Use(static.Serve("/", static.LocalFile("./dist", false)))
  demo.Register(engine)

  container := engine.GetContainer()
  configService := container.MustMake(contract.ConfigKey).(contract.Config)
  if configService.GetBool("app.swagger") == true {
    engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
  }

}
