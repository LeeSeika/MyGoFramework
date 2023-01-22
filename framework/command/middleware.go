package command

import "github.com/godemo/coredemo/framework/cobra"

var middlewareCommand = &cobra.Command{
	Use:   "middleware",
	Short: "迁移中间件的相关命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil
	},
}

var middlewareListCommand = &cobra.Command{
	Use:   "list",
	Short: "列出所有中间件的名称",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
