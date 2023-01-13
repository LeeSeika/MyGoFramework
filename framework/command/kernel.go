package command

import (
	"github.com/godemo/coredemo/framework/cobra"
)

func AddKernelCommand(rootCommand *cobra.Command) {
	rootCommand.AddCommand(DemoCommand)
	rootCommand.AddCommand(initAppCommand())
	rootCommand.AddCommand(initEnvCommand())
	rootCommand.AddCommand(initBuildCommand())
}
