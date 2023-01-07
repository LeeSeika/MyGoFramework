package command

import (
	"fmt"
	"github.com/godemo/coredemo/framework/cobra"
	"github.com/godemo/coredemo/framework/contract"
)

var envCommand = &cobra.Command{
	Use:   "env",
	Short: "获取当前的环境变量",
	Run: func(cmd *cobra.Command, args []string) {
		container := cmd.GetContainer()
		envService := container.MustMake(contract.EnvDevelopment).(contract.Env)
		envMap := envService.AppEnv()
		fmt.Println("环境变量: " + envMap)
	},
}
