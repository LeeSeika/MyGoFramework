package console

import (
	"github.com/godemo/coredemo/app/console/command"
	"github.com/godemo/coredemo/app/console/command/demo"
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/cobra"
)

func RunCommand(container framework.Container) error {
	// 创建根命令
	var rootCommand = &cobra.Command{
		Use:   "hade",
		Short: "hade命令",
		Long:  "hade 框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，也能很方便编写业务命令",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
	rootCommand.SetContainer(container)
	// 加载框架级别命令
	command.AddKernelCommand(rootCommand)
	// 加载业务级别命令
	AddAppCommand(rootCommand)
	return rootCommand.Execute()
}

func AddAppCommand(rootCommand *cobra.Command) {
	rootCommand.AddCommand(demo.InitFoo())
}
