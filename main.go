package main

import (
	"fmt"
	app2 "github.com/godemo/coredemo/app"
	"github.com/godemo/coredemo/app/console"
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/provider/app"
	"github.com/godemo/coredemo/framework/provider/kernel"
)

func main() {
	//core := gin.NewCore()
	//registerRouter(core)
	//server := &http.Server{
	//	Handler: core,
	//	Addr:    ":8081",
	//}
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	fmt.Println("hello")
	//core := gin.New()
	//core.Bind(&demoProvider.DemoProvider{})
	//
	//registerRouter(core)
	//server := &http.Server{
	//	Handler: core,
	//	Addr:    ":8081",
	//}
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	container := framework.NewHadeContainer()
	container.Bind(&app.HadeAppProvider{})
	engine, err := app2.NewHttpEngine()
	if err != nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}
	console.RunCommand(container)
}
