package cmd

import (
	"erbiaoOS/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version 子命令.",
	Long:  "这是一个version 子命令",
	Run:   runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(*cobra.Command, []string) {
	// TODO 这里处理version子命令

	fmt.Println(utils.ProgramName + ": 1.0.0")
}
