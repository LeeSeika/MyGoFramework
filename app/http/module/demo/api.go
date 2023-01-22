package demo

import (
	"fmt"
	demoService "github.com/godemo/coredemo/app/provider/demoProvider"
	"github.com/godemo/coredemo/framework/contract"
	"github.com/godemo/coredemo/framework/gin"
)

type DemoApi struct {
	service *Service
}

func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{service: service}
}

func Register(engine *gin.Engine) error {
	api := NewDemoApi()
	engine.Bind(&demoService.DemoProvider{})
	engine.GET("demo/demo", api.DemoBizApi)
	engine.GET("demo/demo2", api.DemoBizApi2)
	engine.POST("demo/demo_post", api.DemoPostApi)
	engine.GET("demo/demo_config", api.DemoConfigApi)
	return nil
}

func (da *DemoApi) DemoBizApi(context *gin.Context) {
	users := da.service.GetUsers()
	userDTOs := UserModelsToUserDTOs(users)
	context.JSON(200, userDTOs)
	fmt.Println(userDTOs)
	//context.JSON(200, "userDTOs")
}

func (da *DemoApi) DemoBizApi2(context *gin.Context) {
	demoService := context.MustMake(demoService.DemoBizKey).(demoService.IService)
	students := demoService.GetAllStudent()
	userDTOs := StudentsToUserDTOs(students)
	context.JSON(200, userDTOs)
}

func (da *DemoApi) DemoPostApi(context *gin.Context) {
	type Foo struct {
		Name string
	}
	foo := &Foo{}
	err := context.BindJson(foo)
	if err != nil {
		context.AbortWithError(500, err)
	}
	context.JSON(200, nil)
}

func (da *DemoApi) DemoConfigApi(context *gin.Context) {
	configService := context.MustMake(contract.ConfigKey).(contract.Config)
	pwd := configService.GetString("database.mysql.password")
	context.JSON(200, pwd)
}
