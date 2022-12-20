package bizController

import (
	"github.com/godemo/coredemo/app/provider/demoProvider"
	"github.com/godemo/coredemo/framework/gin"
)

func SubjectController(core *gin.Context) {
	demoService := core.MustMake(demoProvider.Key).(*demoProvider.DemoService)
	foo := demoService.GetFoo()
	core.ISetOkStatus().IJson(foo)
}
