package main

//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//)
//
//func GetDemoController(ctx *gin.Context) error {
//	err := ctx.Json(200, "Get/Demo")
//	ctx.Next()
//	return err
//}
//
//func DeleteDemoController(ctx *gin.Context) error {
//	err := ctx.Json(200, "Delete")
//	ctx.Next()
//	return err
//}
//
//func DynamicController(ctx *gin.Context) error {
//	err := ctx.Json(200, "Dynamic")
//	ctx.Next()
//	return err
//}
//
//func CoreFirstController(ctx *gin.Context) error {
//	fmt.Println("CoreFirstController")
//	ctx.Next()
//	return nil
//}
//
//func GroupFirstController(ctx *gin.Context) error {
//	fmt.Println("GroupFirstController")
//	ctx.Next()
//	return nil
//}
//
//func BizMiddleware(ctx *gin.Context) error {
//	fmt.Println("BizMiddleware")
//	ctx.Next()
//	return nil
//}
//
//func Recovery(ctx *gin.Context) error {
//
//	defer func() {
//		if p := recover(); p != nil {
//			ctx.Json(500, "panic")
//		}
//	}()
//
//	ctx.Next()
//
//	return nil
//}
