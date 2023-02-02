package demo

import (
	"database/sql"
	"fmt"
	demoService "github.com/godemo/coredemo/app/provider/demoProvider"
	"github.com/godemo/coredemo/framework/contract"
	"github.com/godemo/coredemo/framework/gin"
	"github.com/godemo/coredemo/framework/provider/orm"
	"time"
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
	engine.GET("demo/demo_orm", api.DemoOrm)
	return nil
}

func (da *DemoApi) DemoBizApi(context *gin.Context) {
	users := da.service.GetUsers()
	userDTOs := UserModelsToUserDTOs(users)
	context.JSON(200, userDTOs)
	fmt.Println(userDTOs)
	//context.JSON(200, "userDTOs")
}

// DemoBizApi2 Demo2 for godoc
// @Summary 获取所有学生
// @Description 获取所有学生，不进行分页
// @Produce json
// @Tags demo
// @Success 200 {array} []UserDTO
// @Router /demo/demo2 [get]
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

func (da *DemoApi) DemoOrm(context *gin.Context) {
	logService := context.MustMake(contract.LogKey).(contract.Log)
	logService.Info(context, "request start", nil)

	ormService := context.MustMake(contract.ORMKey).(contract.ORM)
	db, err := ormService.GetDB(orm.WithConfigPath("database.default"))
	if err != nil {
		logService.Error(context, err.Error(), nil)
		context.AbortWithError(500, err)
		return
	}
	db.WithContext(context)

	err = db.AutoMigrate(&UserForGorm{})
	if err != nil {
		context.AbortWithError(500, err)
		return
	}
	logService.Info(context, "migrate ok", nil)

	email := "additcd@qq.com"
	name := "leeseika"
	age := uint8(24)
	brithday := time.Date(1999, 2, 25, 1, 1, 1, 1, time.Local)
	user := &UserForGorm{
		Name:         name,
		Email:        &email,
		Age:          age,
		Birthday:     &brithday,
		MemberNumber: sql.NullString{},
		ActivatedAt:  sql.NullTime{},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err = db.Create(user).Error
	logService.Info(context, "insert user", map[string]interface{}{
		"err": err,
		"id":  user.ID,
	})

	user.Name = "update"
	err = db.Save(user).Error
	logService.Info(context, "update user", map[string]interface{}{
		"err": err,
		"id":  user.ID,
	})

	queryUser := &UserForGorm{ID: user.ID}
	err = db.First(queryUser).Error
	logService.Info(context, "query user", map[string]interface{}{
		"err": err,
		"id":  user.ID,
	})

	err = db.Delete(queryUser).Error
	logService.Info(context, "delete user", map[string]interface{}{
		"err": err,
		"id":  user.ID,
	})
	context.JSON(200, "ok")
}
