package command

import (
	"fmt"
	"github.com/godemo/coredemo/framework/cobra"
	"github.com/godemo/coredemo/framework/contract"
)

var DemoCommand = &cobra.Command{
	Use:   "demo",
	Short: "demo 命令",
	Run: func(cmd *cobra.Command, args []string) {
		container := cmd.GetContainer()
		app := container.MustMake(contract.AppKey).(contract.App)
		fmt.Println("base folder:" + app.BaseFolder())
	},
}
