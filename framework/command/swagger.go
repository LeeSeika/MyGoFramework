package command

import (
	"fmt"
	"github.com/godemo/coredemo/framework/cobra"
	"github.com/godemo/coredemo/framework/contract"
	"github.com/swaggo/swag/gen"
	"path/filepath"
)

func initSwaggerCommand() *cobra.Command {
	swaggerCommand.AddCommand(swaggerGenCommand)
	return swaggerCommand
}

var swaggerCommand = &cobra.Command{
	Use:   "swagger",
	Short: "swagger相关命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	},
}

var swaggerGenCommand = &cobra.Command{
	Use:   "gen",
	Short: "生成对应的swagger文件, contain swagger.yaml, doc.go",
	Run: func(cmd *cobra.Command, args []string) {
		container := cmd.GetContainer()
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
			OutputTypes:        []string{"go", "json", "yaml"},
		}
		err := gen.New().Build(conf)
		if err != nil {
			fmt.Println(err)
		}
	},
}
