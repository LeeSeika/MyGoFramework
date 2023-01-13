package command

import (
	"fmt"
	"github.com/godemo/coredemo/framework/cobra"
	"log"
	"os/exec"
)

var buildCommand = &cobra.Command{
	Use:   "build",
	Short: "编译相关命令",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

var buildSelfCommand = &cobra.Command{
	Use:   "",
	Short: "",
	RunE: func(c *cobra.Command, args []string) error {
		path, err := exec.LookPath("go")
		if err != nil {
			log.Fatalln("coredemo go: please install go in path first")
		}

		cmd := exec.Command(path, "build", "-o", "coredemo", "./")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("go build error:")
			fmt.Println(string(out))
			fmt.Println("--------------")
			return err
		}
		fmt.Println("build success please run ./coredemo direct")
		return nil
	},
}

var buildFrontendCommand = &cobra.Command{
	Use:   "frontend",
	Short: "使用npm编译前端",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := exec.LookPath("npm")
		if err != nil {
			log.Fatalln("请安装npm在PATH目录下")
		}
		command := exec.Command(path, "run", "build")
		output, err := command.CombinedOutput()
		if err != nil {
			fmt.Println("============= 前端编译失败 ============")
			fmt.Println(string(output))
			fmt.Println("============= 前端编译失败 ============")
			return err
		}
		fmt.Println(string(output))
		fmt.Println("============= 前端编译成功 ============")
		return nil
	},
}

var buildBackendCommand = &cobra.Command{
	Use:   "backend",
	Short: "编译go后端",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := buildSelfCommand.RunE(cmd, args)
		return err
	},
}

var buildAllCommand = &cobra.Command{
	Use:   "all",
	Short: "编译前端和后端",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := buildFrontendCommand.RunE(cmd, args)
		if err != nil {
			return err
		}
		err = buildBackendCommand.RunE(cmd, args)
		if err != nil {
			return err
		}
		return nil
	},
}

func initBuildCommand() *cobra.Command {
	buildCommand.AddCommand(buildFrontendCommand)
	buildCommand.AddCommand(buildBackendCommand)
	buildCommand.AddCommand(buildAllCommand)
	return buildCommand
}
