package main

import (
	"context"
	"fmt"
	app2 "github.com/godemo/coredemo/app"
	"github.com/godemo/coredemo/app/console"
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/contract"
	"github.com/godemo/coredemo/framework/provider/app"
	"github.com/godemo/coredemo/framework/provider/config"
	"github.com/godemo/coredemo/framework/provider/env"
	"github.com/godemo/coredemo/framework/provider/kernel"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	container.Bind(&env.HadeEnvProvider{})
	engine, err := app2.NewHttpEngine()
	if err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}
	container.Bind(&config.HadeConfigProvider{})

	console.RunCommand(container)
	// startServer(container)
	//proxy := command.NewProxy(container)
	//go proxy.MonitorBackend()
	//if err := proxy.StartProxy(true, true); err != nil {
	//
	//}
}

func startServer(container framework.Container) {
	// 从服务容器中获取kernel的服务实例
	kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)
	// 从kernel服务实例中获取引擎
	core := kernelService.HttpEngine()

	// 创建一个Server服务
	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}

	// 这个goroutine是启动服务的goroutine
	go func() {
		server.ListenAndServe()
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
