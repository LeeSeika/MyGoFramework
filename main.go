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
	"github.com/swaggo/swag/gen"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func main() {

	container := framework.NewHadeContainer()
	container.Bind(&app.HadeAppProvider{})
	container.Bind(&env.HadeEnvProvider{})
	container.Bind(&config.HadeConfigProvider{})
	engine, err := app2.NewHttpEngine(container)
	if err == nil {
		container.Bind(&kernel.HadeKernelProvider{HttpEngine: engine})
	}

	console.RunCommand(container)
	//testSwagger(container)
}

func testSwagger(container framework.Container) {
	appService := container.MustMake(contract.AppKey).(contract.App)
	outputDir := filepath.Join(appService.AppFolder(), "http", "swagger")
	httpDir := filepath.Join(appService.AppFolder(), "http")
	conf := &gen.Config{
		SearchDir:          httpDir,
		Excludes:           "",
		MainAPIFile:        "swagger.go",
		OutputDir:          outputDir,
		PropNamingStrategy: "",
		ParseVendor:        false,
		ParseDependency:    false,
		ParseInternal:      false,
		MarkdownFilesDir:   "",
		GeneratedTime:      false,
		OutputTypes:        []string{"json", "yaml", "go"},
	}
	err := gen.New().Build(conf)
	if err != nil {
		fmt.Println(err)
	}
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
