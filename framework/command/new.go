package command

import (
	"bytes"
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/godemo/coredemo/framework/cobra"
	"github.com/godemo/coredemo/framework/util"
	"github.com/google/go-github/v39/github"
	"github.com/spf13/cast"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func initNewCommand() *cobra.Command {
	return newCommand
}

var newCommand = &cobra.Command{
	Use:   "new",
	Short: "新建一个示例项目",
	RunE: func(cmd *cobra.Command, args []string) error {
		currentPath := util.GetExecDirectory()

		var name string
		var folder string
		var mod string
		var version string
		var release *github.RepositoryRelease
		{
			prompt := &survey.Input{
				Message: "请输入目录名称：",
			}
			err := survey.AskOne(prompt, &name)
			if err != nil {
				return err
			}

			folder = filepath.Join(currentPath, name)
			if util.Exists(folder) {
				force := false
				prompt2 := &survey.Confirm{
					Message: "已存在重复目录名称，是否删除该目录？",
					Default: false,
				}
				err := survey.AskOne(prompt2, &force)
				if err != nil {
					return err
				}
				if force {
					err := os.RemoveAll(folder)
					if err != nil {
						return err
					}
				} else {
					fmt.Println("目录已存在，创建失败")
					return nil
				}
			}
		}
		{
			prompt := &survey.Input{
				Message: "请输入go.mod中module的名称，默认为目录名称",
			}
			err := survey.AskOne(prompt, &mod)
			if err != nil {
				return err
			}
			if mod == "" {
				mod = folder
			}
		}
		{
			client := github.NewClient(nil)
			prompt := &survey.Input{
				Message: "请输入版本号（默认为最新版本）",
			}
			err := survey.AskOne(prompt, &version)
			if err != nil {
				return err
			}
			if version != "" {
				release, _, err = client.Repositories.GetReleaseByTag(context.Background(), "LeeSeika", "MyGoFramework", version)
				if err != nil && release == nil {
					fmt.Println("版本不存在")
					return nil
				}
			}
			if version == "" {
				release, _, err = client.Repositories.GetLatestRelease(context.Background(), "LeeSeika", "MyGoFramework")
				version = release.GetTagName()
			}
		}
		fmt.Println("====================================================")
		fmt.Println("开始进行创建应用操作")
		fmt.Println("创建目录：", folder)
		fmt.Println("应用名称：", mod)
		fmt.Println("coredemo框架版本：", release.GetTagName())

		templateFolder := filepath.Join(currentPath, "template-coredemo-"+version+"-"+cast.ToString(time.Now().Unix()))
		os.Mkdir(templateFolder, os.ModePerm)
		fmt.Println("创建临时目录", templateFolder)

		url := release.GetZipballURL()
		err := util.DownloadFile(filepath.Join(templateFolder, "template.zip"), url)
		if err != nil {
			return err
		}
		fmt.Println("已下载zip包")
		_, err = util.Unzip(filepath.Join(templateFolder, "template.zip"), templateFolder)
		if err != nil {
			return err
		}
		fileInfos, err := ioutil.ReadDir(templateFolder)
		if err != nil {
			return err
		}
		for _, fileInfo := range fileInfos {
			if fileInfo.IsDir() && strings.Contains(fileInfo.Name(), "LeeSeika-MyGoFramework-") {
				if err := os.Rename(filepath.Join(templateFolder, fileInfo.Name()), folder); err != nil {
					return err
				}
			}
		}
		fmt.Println("解压zip包")

		if err := os.RemoveAll(templateFolder); err != nil {
			return err
		}
		fmt.Println("删除临时文件夹")

		os.RemoveAll(path.Join(folder, ".git"))
		fmt.Println("删除.git目录")

		// 删除framework 目录
		os.RemoveAll(path.Join(folder, "framework"))
		fmt.Println("删除framework目录")

		filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if info != nil && info.IsDir() {
				return nil
			}

			c, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			if path == filepath.Join(folder, "go.mod") {
				fmt.Println("更新文件:" + path)
				c = bytes.ReplaceAll(c, []byte("module github.com/godemo/coredemo"), []byte("module "+mod))
				c = bytes.ReplaceAll(c, []byte("require ("), []byte("require (\n\tgithub.com/godemo/coredemo "+version))
				err = ioutil.WriteFile(path, c, 0644)
				if err != nil {
					return err
				}
				return nil
			}

			isContain := bytes.Contains(c, []byte("github.com/godemo/coredemo/app"))
			if isContain {
				fmt.Println("更新文件:" + path)
				c = bytes.ReplaceAll(c, []byte("github.com/godemo/coredemo/app"), []byte(mod+"/app"))
				err = ioutil.WriteFile(path, c, 0644)
				if err != nil {
					return err
				}
			}

			return nil
		})
		fmt.Println("创建应用结束")
		fmt.Println("目录：", folder)
		fmt.Println("====================================================")
		return nil
	},
}
