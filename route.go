package main

import (
	"github.com/godemo/coredemo/bizController"
	"github.com/godemo/coredemo/framework/gin"
)

//import "coredemo/framework"
//
//func registerRouter(core *framework.Core) {
//	core.Use(CoreFirstController)
//	core.GetDynamic("/user/getDemo", GetDemoController, BizMiddleware)
//
//	group := framework.NewGroup(core, "/prefix")
//	{
//		group.Use(GroupFirstController)
//		group.GetDynamic("/dynamic/:id", DynamicController)
//		group.PostDynamic("/dynamic/:id", DynamicController)
//		group.DeleteDynamic("/deleteDemo", DeleteDemoController)
//	}
//}

func registerRouter(core *gin.Engine) {
	core.GET("/subject/foo", bizController.SubjectController)
}
