package command

import (
	"context"
	"errors"
	"fmt"
	"github.com/erikdubbelboer/gspt"
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/cobra"
	"github.com/godemo/coredemo/framework/contract"
	"github.com/godemo/coredemo/framework/util"
	"github.com/sevlyar/go-daemon"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

var appAddress = ""
var appDaemon = false

// initAppCommand 初始化app命令和其子命令
func initAppCommand() *cobra.Command {
	appStartCommand.Flags().BoolVarP(&appDaemon, "daemon", "d", false, "start app daemon")
	appStartCommand.Flags().StringVar(&appAddress, "address", ":8080", "设置app启动的地址，默认为:8888")
	appCommand.AddCommand(appStartCommand)
	appCommand.AddCommand(appStopCommand)
	appCommand.AddCommand(appStateCommand)
	appCommand.AddCommand(appReStartCommand)
	return appCommand
}

// AppCommand 是命令行参数第一级为app的命令，它没有实际功能，只是打印帮助文档
var appCommand = &cobra.Command{
	Use:   "app",
	Short: "业务应用控制命令",
	Long:  "业务应用控制命令，其包含业务启动，关闭，重启，查询等功能",
	RunE: func(c *cobra.Command, args []string) error {
		// 打印帮助文档
		c.Help()
		return nil
	},
}

// appStartCommand 启动一个Web服务
var appStartCommand = &cobra.Command{
	Use:   "start",
	Short: "启动一个Web服务",
	RunE: func(c *cobra.Command, args []string) error {
		// 从Command中获取服务容器
		container := c.GetContainer()
		// 从服务容器中获取kernel的服务实例
		kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)
		// 从kernel服务实例中获取引擎
		core := kernelService.HttpEngine()

		if appAddress == "" {
			envService := container.MustMake(contract.EnvKey).(contract.Env)
			if envService.IsExist("ADDRESS") {
				appAddress = envService.Get("ADDRESS")
			} else {
				configService := container.MustMake(contract.ConfigKey).(contract.Config)
				if configService.IsExist("app.address") {
					appAddress = configService.GetString("app.address")
				} else {
					appAddress = ":8080"
				}
			}
		}

		// 创建一个Server服务
		server := &http.Server{
			Handler: core,
			Addr:    appAddress,
		}

		// check daemon
		appService := container.MustMake(contract.AppKey).(contract.App)
		pidFolder := appService.PidFolder()

		if !util.Exists(pidFolder) {
			if err := os.MkdirAll(pidFolder, os.ModePerm); err != nil {
				return err
			}
		}
		pidFilePath := filepath.Join(pidFolder, "app.pid")

		logFolder := appService.LogFolder()
		if !util.Exists(logFolder) {
			if err := os.MkdirAll(logFolder, os.ModePerm); err != nil {
				return err
			}
		}
		logFilePath := filepath.Join(logFolder, "app.log")
		currentFolder := util.GetExecDirectory()

		// daemon
		if appDaemon {
			daemonCtx := &daemon.Context{
				PidFileName: pidFilePath,
				PidFilePerm: 0664,
				LogFileName: logFilePath,
				LogFilePerm: 0640,
				WorkDir:     currentFolder,
				Umask:       027,
				Args:        []string{"", "app", "start", "--daemon=true"},
			}
			child, err := daemonCtx.Reborn()
			if err != nil {
				return err
			}
			if child != nil {
				fmt.Println("app启动成功, pid:", child.Pid)
				fmt.Println("日志文件：", logFilePath)
				return nil
			}
			defer daemonCtx.Release()
			fmt.Println("daemon started")
			gspt.SetProcTitle("coredemo app")
			if err := startAppServe(server, container); err != nil {
				return err
			}
			return nil
		}

		// not daemon
		content := strconv.Itoa(os.Getpid())
		fmt.Println("[PID]", content)
		err := ioutil.WriteFile(pidFilePath, []byte(content), 0644)
		if err != nil {
			return err
		}
		gspt.SetProcTitle("coredemo app")

		fmt.Println("app serve url:", appAddress)
		err = startAppServe(server, container)
		return err
	},
}

var appStateCommand = &cobra.Command{
	Use:   "state",
	Short: "获取进程状态",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)
		pidFolder := appService.PidFolder()
		pidFilePath := filepath.Join(pidFolder, "app.pid")
		content, err := ioutil.ReadFile(pidFilePath)
		if err != nil {
			return err
		}
		if content == nil || len(content) < 1 {
			fmt.Println("没有启动app服务")
			return nil
		}
		pid, err := strconv.Atoi(string(content))
		if err != nil {
			return err
		}
		exist := util.CheckProcessExist(pid)
		if exist {
			fmt.Println("app服务已经启动, pid:", pid)
			return nil
		}
		fmt.Println("没有启动app服务")
		return nil
	},
}

var appStopCommand = &cobra.Command{
	Use:   "stop",
	Short: "停止app服务进程",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)
		pidFolder := appService.PidFolder()
		pidFilePath := filepath.Join(pidFolder, "app.pid")
		content, err := ioutil.ReadFile(pidFilePath)
		if err != nil {
			return err
		}
		if content != nil && len(content) > 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				return err
			}
			if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
				return err
			}
			if err := ioutil.WriteFile(pidFilePath, []byte{}, 0644); err != nil {
				return err
			}
			fmt.Println("停止进程:", pid)
		}
		return nil
	},
}

var appReStartCommand = &cobra.Command{
	Use:   "restart",
	Short: "重启app服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)
		pidFolder := appService.PidFolder()
		pidFilePath := filepath.Join(pidFolder, "app.id")
		content, err := ioutil.ReadFile(pidFilePath)
		if err != nil {
			return err
		}
		if content != nil && len(content) > 0 {
			pid, err := strconv.Atoi(string(content))
			if err != nil {
				return err
			}
			exist := util.CheckProcessExist(pid)
			if exist {
				err := syscall.Kill(pid, syscall.SIGTERM)
				if err != nil {
					return err
				}
				err = ioutil.WriteFile(pidFilePath, []byte{}, 0644)
				if err != nil {
					return err
				}
				closeWait := 5
				configService := container.MustMake(contract.ConfigKey).(contract.Config)
				if configService.IsExist("app.close_wait") {
					closeWait = configService.GetInt("app.close_wait")
				}
				for i := 0; i < closeWait*2; i++ {
					if util.CheckProcessExist(pid) == false {
						break
					}
					time.Sleep(1 * time.Second)
				}
				if util.CheckProcessExist(pid) == true {
					fmt.Println("结束进程失败:"+strconv.Itoa(pid), "请查看原因")
					return errors.New("结束进程失败")
				}
				fmt.Println("进程结束成功")
			}

		}
		appDaemon = true
		return appStartCommand.RunE(cmd, args)
	},
}

func startAppServe(server *http.Server, container framework.Container) error {
	// 这个goroutine是启动服务的goroutine
	go func() {
		server.ListenAndServe()
	}()

	closeWait := 5
	configService := container.MustMake(contract.ConfigKey).(contract.Config)
	if configService.IsExist("app.close_wait") {
		closeWait = configService.GetInt("app.close_wait")
	}

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	// 调用Server.Shutdown graceful结束
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Duration(closeWait)*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	return nil
}
