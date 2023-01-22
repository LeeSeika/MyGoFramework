package test
import (
  "fmt"
  "github.com/godemo/coredemo/framework/cobra"
)
var TestCommand = &cobra.Command {
  Use: "test",
  Short: "test",
  RunE: func(cmd *cobra.Command, args []string) error {
    ctn := cmd.GetContainer()
    fmt.Println(ctn)
    return nil
  },
}
