package command

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/godemo/coredemo/framework/cobra"
	"github.com/godemo/coredemo/framework/contract"
	"github.com/godemo/coredemo/framework/util"
	"github.com/pkg/errors"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var cmdCommand = &cobra.Command{
	Use:   "command",
	Short: "自定义命令行工具",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil
	},
}

var cmdListCommand = &cobra.Command{
	Use:   "list",
	Short: "查看所有命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		allCommands := cmd.Root().Commands()
		for _, command := range allCommands {
			fmt.Println(command.Name() + ": " + command.Short)
		}
		return nil
	},
}

var cmdNewCommand = &cobra.Command{
	Use:   "new",
	Short: "创建一个新的命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name string
		var folder string
		{
			prompt := &survey.Input{
				Message: "请输入命令的名称",
			}
			err := survey.AskOne(prompt, &name)
			if err != nil {
				return err
			}
		}
		{
			prompt := &survey.Input{
				Message: "请输入文件夹名称（同命令名称）",
			}
			err := survey.AskOne(prompt, &folder)
			if err != nil {
				return err
			}
		}
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)
		commandFolder := appService.CommandFolder()
		// 检查文件夹是否重复
		subDirs, err := util.SubDir(commandFolder)
		if err != nil {
			return errors.Cause(err)
		}
		for _, subDir := range subDirs {
			if subDir == folder {
				fmt.Println("目录" + "\"" + folder + "\"已经存在")
				return nil
			}
		}
		err = os.Mkdir(filepath.Join(commandFolder, folder), 0700)
		if err != nil {
			return err
		}
		fullFilePath := filepath.Join(commandFolder, folder, name+".go")
		file, err := os.Create(fullFilePath)
		if err != nil {
			return errors.Cause(err)
		}
		funcs := template.FuncMap{"title": strings.Title}
		t := template.Must(template.New("cmd").Funcs(funcs).Parse(cmdTmpl))
		err = t.Execute(file, name)
		if err != nil {
			return errors.Cause(err)
		}
		return nil
	},
}

func initCmdCommands() *cobra.Command {
	cmdCommand.AddCommand(cmdNewCommand)
	cmdCommand.AddCommand(cmdListCommand)
	return cmdCommand
}

var cmdTmpl string = `package {{.}}
import (
  "fmt"
  "github.com/godemo/coredemo/framework/cobra"
)
var {{.|title}}Command = &cobra.Command {
  Use: "{{.}}",
  Short: "{{.}}",
  RunE: func(cmd *cobra.Command, args []string) error {
    ctn := cmd.GetContainer()
    fmt.Println(ctn)
    return nil
  },
}
`
