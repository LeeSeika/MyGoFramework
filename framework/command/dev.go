package command

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/cobra"
	"github.com/godemo/coredemo/framework/contract"
	"github.com/godemo/coredemo/framework/util"
	"io/fs"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

type Proxy struct {
	devConfig   *devConfig
	frontendPid int
	backendPid  int
}

var devCommand = &cobra.Command{
	Use:   "dev",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	},
}

var devFrontendCommand = &cobra.Command{
	Use:   "frontend",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		proxy := NewProxy(cmd.GetContainer())
		err := proxy.StartProxy(true, false)
		return err
	},
}

var devBackendCommand = &cobra.Command{
	Use:   "backend",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		proxy := NewProxy(cmd.GetContainer())
		go proxy.MonitorBackend()
		if err := proxy.StartProxy(false, true); err != nil {
			fmt.Println(err.Error())
			return err
		}
		return nil
	},
}

var devAllCommand = &cobra.Command{
	Use:   "all",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		proxy := NewProxy(cmd.GetContainer())
		go proxy.MonitorBackend()
		if err := proxy.StartProxy(true, true); err != nil {
			return err
		}
		return nil
	},
}

type devConfig struct {
	Port    string
	Backend struct {
		RefreshTime   int
		Port          string
		MonitorFolder string
	}

	Frontend struct {
		Port string
	}
}

func initDevCommand() *cobra.Command {
	devCommand.AddCommand(devFrontendCommand)
	devCommand.AddCommand(devBackendCommand)
	devCommand.AddCommand(devAllCommand)
	return devCommand
}

func initDevConfig(container framework.Container) *devConfig {
	devConfig := &devConfig{
		Port: "8070",
		Backend: struct {
			RefreshTime   int
			Port          string
			MonitorFolder string
		}{
			1,
			"8072",
			"",
		},
		Frontend: struct {
			Port string
		}{
			"8071",
		},
	}
	configService := container.MustMake(contract.ConfigKey).(contract.Config)
	if configService.IsExist("app.dev.port") {
		port := configService.GetString("app.dev.port")
		devConfig.Port = port
	}
	if configService.IsExist("app.dev.backend.port") {
		port := configService.GetString("app.dev.backend.port")
		devConfig.Backend.Port = port
	}
	if configService.IsExist("app.dev.backend.refresh_time") {
		refreshTime := configService.GetInt("app.dev.backend.refresh_time")
		devConfig.Backend.RefreshTime = refreshTime
	}
	monitorFolder := configService.GetString("app.dev.backend.monitor_folder")
	if monitorFolder == "" {
		appService := container.MustMake(contract.AppKey).(contract.App)
		monitorFolder = appService.AppFolder()
	}
	devConfig.Backend.MonitorFolder = monitorFolder
	if configService.IsExist("app.dev.frontend.port") {
		port := configService.GetString("app.dev.frontend.port")
		devConfig.Frontend.Port = port
	}

	return devConfig
}

func NewProxy(c framework.Container) *Proxy {
	devConfig := initDevConfig(c)
	return &Proxy{devConfig: devConfig}
}

func (p *Proxy) newProxyReverseProxy(frontend, backend *url.URL) *httputil.ReverseProxy {
	if p.frontendPid == 0 && p.backendPid == 0 {
		fmt.Println("前端后端服务都不存在")
		return nil
	}
	if p.frontendPid == 0 && p.backendPid != 0 {
		return httputil.NewSingleHostReverseProxy(backend)
	}
	if p.frontendPid != 0 && p.backendPid == 0 {
		return httputil.NewSingleHostReverseProxy(frontend)
	}

	director := func(req *http.Request) {
		if req.URL.Path == "/" || req.URL.Path == "/app.js" {
			req.URL.Scheme = frontend.Scheme
			req.URL.Host = frontend.Host
		} else {
			req.URL.Scheme = backend.Scheme
			req.URL.Host = backend.Host
		}
	}

	NotFoundErr := errors.New("response is 404, need to redirect")
	return &httputil.ReverseProxy{
		Director: director,
		ModifyResponse: func(response *http.Response) error {
			if response.StatusCode == 404 {
				return NotFoundErr
			}
			return nil
		},
		ErrorHandler: func(writer http.ResponseWriter, request *http.Request, err error) {
			if errors.Is(err, NotFoundErr) {
				httputil.NewSingleHostReverseProxy(frontend).ServeHTTP(writer, request)
			}
		},
	}
}

func (p *Proxy) StartProxy(startFrontend, startBackend bool) error {
	var backendURL, frontendURL *url.URL
	var err error
	if startBackend {
		if err := p.restartBackend(); err != nil {
			return err
		}
	}
	if startFrontend {
		if err := p.restartFrontend(); err != nil {
			return err
		}
	}
	if frontendURL, err = url.Parse(fmt.Sprintf("%s%s", "http://127.0.0.1:", p.devConfig.Frontend.Port)); err != nil {
		return err
	}
	if backendURL, err = url.Parse(fmt.Sprintf("%s%s", "http://localhost:", p.devConfig.Backend.Port)); err != nil {
		return err
	}

	proxyReverse := p.newProxyReverseProxy(frontendURL, backendURL)
	proxyServer := &http.Server{
		Addr:    "127.0.0.1:" + p.devConfig.Port,
		Handler: proxyReverse,
	}
	fmt.Println("代理服务启动:", "http://"+proxyServer.Addr)
	err = proxyServer.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (p *Proxy) restartFrontend() error {
	if p.frontendPid != 0 {
		syscall.Kill(p.frontendPid, syscall.SIGKILL)
		p.frontendPid = 0
	}
	port := p.devConfig.Frontend.Port
	path, err := exec.LookPath("npm")
	if err != nil {
		return err
	}
	cmd := exec.Command(path, "run", "dev")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s%s", "PORT=", port))
	cmd.Stdout = os.NewFile(0, os.DevNull)
	cmd.Stderr = os.Stderr
	fmt.Println(cmd.Env)
	err = cmd.Start()
	fmt.Println("启动前端服务", "http://127.0.0.1:"+port)
	if err != nil {
		fmt.Println(err)
	}
	p.frontendPid = cmd.Process.Pid
	fmt.Println("前端pid:", p.frontendPid)
	return nil
}

func (p *Proxy) restartBackend() error {
	if p.backendPid != 0 {
		syscall.Kill(p.backendPid, syscall.SIGKILL)
		p.backendPid = 0
	}
	port := p.devConfig.Backend.Port
	addr := fmt.Sprintf(":" + port)
	cmd := exec.Command("./coredemo", "app", "start", "--address="+addr)
	cmd.Stdout = os.NewFile(0, os.DevNull)
	cmd.Stderr = os.Stderr
	fmt.Println("启动后端服务: ", "http://localhost:"+port)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	p.backendPid = cmd.Process.Pid
	fmt.Println("后端服务pid: ", p.backendPid)
	return nil
}

func (p *Proxy) rebuildBackend() error {
	cmdBuild := exec.Command("./coredemo", "build", "backend")
	cmdBuild.Stdout = os.Stdout
	cmdBuild.Stderr = os.Stderr
	if err := cmdBuild.Start(); err != nil {
		err := cmdBuild.Wait()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Proxy) MonitorBackend() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	defer watcher.Close()
	monitorFolder := p.devConfig.Backend.MonitorFolder
	fmt.Println("监控文件夹" + monitorFolder)
	filepath.Walk(monitorFolder, func(path string, info fs.FileInfo, err error) error {
		if info != nil && !info.IsDir() {
			return nil
		}
		if util.IsHiddenDirectory(path) {
			return nil
		}
		fmt.Println(path)
		return watcher.Add(path)
	})
	refreshTime := p.devConfig.Backend.RefreshTime
	timer := time.NewTimer(time.Duration(refreshTime) * time.Second)
	timer.Stop()
	for {
		select {
		case <-timer.C:
			fmt.Println("...检测到文件改动，重新编译")
			if err := p.rebuildBackend(); err != nil {
				fmt.Println("编译失败:", err.Error())
			} else {
				if err := p.restartBackend(); err != nil {
					fmt.Println("重启服务失败:", err.Error())
				}
			}
		case _, ok := <-watcher.Events:
			if !ok {
				continue
			}
			timer.Reset(time.Duration(refreshTime) * time.Second)
		case err, ok := <-watcher.Errors:
			if !ok {
				continue
			}
			fmt.Println("监听文件夹发生错误：", err.Error())
			timer.Reset(time.Duration(refreshTime) * time.Second)
		}
	}
}

func (p *Proxy) NewReverseProxy(frontend, backend *url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.URL.Scheme = backend.Scheme
		req.URL.Host = backend.Host
	}
	NotFoundErr := errors.New("response is 404, need to redirect")
	return &httputil.ReverseProxy{
		Director: director,
		ModifyResponse: func(response *http.Response) error {
			if response.StatusCode == 404 {
				return NotFoundErr
			}
			return nil
		},
		ErrorHandler: func(writer http.ResponseWriter, request *http.Request, err error) {
			if errors.Is(err, NotFoundErr) {
				httputil.NewSingleHostReverseProxy(frontend).ServeHTTP(writer, request)
			}
		},
	}
}
