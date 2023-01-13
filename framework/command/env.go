package command

import (
	"fmt"
	"github.com/godemo/coredemo/framework/cobra"
	"github.com/godemo/coredemo/framework/contract"
)

var envCommand = &cobra.Command{
	Use:   "env",
	Short: "获取当前的环境变量APP_ENV",
	Run: func(cmd *cobra.Command, args []string) {
		container := cmd.GetContainer()
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		appEnv := envService.AppEnv()
		fmt.Println("环境变量APP_ENV: " + appEnv)
	},
}

var envListCommand = &cobra.Command{
	Use:   "list",
	Short: "获取所有环境变量",
	Run: func(cmd *cobra.Command, args []string) {
		container := cmd.GetContainer()
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		allEnv := envService.All()
		fmt.Println("全部环境变量:")
		for k, v := range allEnv {
			fmt.Println("key: " + k + ", value: " + v)
		}
	},
}

func initEnvCommand() *cobra.Command {
	envCommand.AddCommand(envListCommand)
	return envCommand
}
