package command

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/cobra"
	"github.com/godemo/coredemo/framework/contract"
	"github.com/godemo/coredemo/framework/util"
	"github.com/jianfengye/collection"
	"github.com/pkg/errors"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var providerCommand = &cobra.Command{
	Use:   "provider",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}
		return nil
	},
}

var providerListCommand = &cobra.Command{
	Use:   "list",
	Short: "列出容器内的所有服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer().(*framework.HadeContainer)
		nameList := container.NameList()
		for _, line := range nameList {
			println(line)
		}
		return nil
	},
}

var providerCreateCommand = &cobra.Command{
	Use:     "new",
	Short:   "创建一个服务",
	Aliases: []string{"create", "init"},
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		fmt.Println("创建一个服务")
		var name string
		var folder string
		{
			prompt := &survey.Input{
				Message: "请输入服务名称（服务凭证）",
			}
			err := survey.AskOne(prompt, &name)
			if err != nil {
				return err
			}
		}
		{
			prompt := &survey.Input{
				Message: "请输入服务所在的目录名称（跟服务名称相同）",
			}
			err := survey.AskOne(prompt, &folder)
			if err != nil {
				return err
			}
		}
		providerList := container.(*framework.HadeContainer).NameList()
		listCollection := collection.NewStrCollection(providerList)
		if listCollection.Contains(name) {
			fmt.Println("服务\"" + name + "\"已存在")
			return nil
		}
		if folder == "" {
			folder = name
		}
		appService := container.MustMake(contract.AppKey).(contract.App)
		providerFolder := appService.ProviderFolder()
		subFolders, err := util.SubDir(providerFolder)
		if err != nil {
			return err
		}
		subFolderCollection := collection.NewStrCollection(subFolders)
		if subFolderCollection.Contains(folder) {
			fmt.Println("已存在名为\"" + folder + "\"的目录")
			return nil
		}
		if err = os.Mkdir(filepath.Join(providerFolder, folder), 0700); err != nil {
			return err
		}
		funcs := template.FuncMap{"title": strings.Title}
		{
			// contract.go
			fullFilePath := filepath.Join(providerFolder, folder, "contract.go")
			file, err := os.Create(fullFilePath)
			if err != nil {
				return errors.Cause(err)
			}
			t := template.Must(template.New("contract").Funcs(funcs).Parse(contractTmp))
			if err := t.Execute(file, name); err != nil {
				return errors.Cause(err)
			}

		}
		{
			// provider.go
			fullFilePath := filepath.Join(providerFolder, folder, "provider.go")
			file, err := os.Create(fullFilePath)
			if err != nil {
				return errors.Cause(err)
			}
			t := template.Must(template.New("provider").Funcs(funcs).Parse(providerTmp))
			if err := t.Execute(file, name); err != nil {
				return errors.Cause(err)
			}
		}
		{
			// service.go
			fullFilePath := filepath.Join(providerFolder, folder, name+".go")
			file, err := os.Create(fullFilePath)
			if err != nil {
				return errors.Cause(err)
			}
			t := template.Must(template.New("service").Funcs(funcs).Parse(serviceTmp))
			if err := t.Execute(file, name); err != nil {
				return errors.Cause(err)
			}
			fmt.Println("创建服务成功，文件夹地址：", filepath.Join(providerFolder, folder))
			return nil
		}
	},
}

var contractTmp string = `package {{.}}
const {{.|title}}Key = "{{.}}"
type Service interface {
  Foo() string
}
`

var providerTmp string = `package {{.}}
import (
  "github.com/godemo/coredemo/framework"
)
type {{.|title}}Provider struct {
  framework.ServiceProvider
  c framework.Container
}

func (sp *{{.|title}}Provider) Name() string {
  return {{.|title}}Key
}

func (sp *{{.|title}}Provider) Register(c framework.Container) framework.NewInstance {
  return New{{.|title}}Service
}

func (sp *{{.|title}}Provider) IsDefer() bool {
  return false
}

func (sp *{{.|title}}Provider) Params(c framework.Container) []interface{} {
  return []interface{}{c}
}

func (sp *{{.|title}}Provider) Boot(c framework.Container) error {
  return nil
}
`

var serviceTmp string = `package {{.}}
import (
  "github.com/godemo/coredemo/framework"
)

type {{.|title}}Service struct {
  container framework.Container
}

func New{{.|title}}Service(params ...interface{}) (interface{}, error) {
  container := params[0].(framework.Container)
  return &{{.|title}}Service{container: container}, nil
}

func (s *{{.|title}}Service) Foo() string {
  return ""
}
`

func initProviderCommand() *cobra.Command {
	providerCommand.AddCommand(providerListCommand)
	providerCommand.AddCommand(providerCreateCommand)
	return providerCommand
}
