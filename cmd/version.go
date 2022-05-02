package cmd

import (
	"erbiaoOS/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "查询版本",
	Long:  fmt.Sprintf("查询当前%s版本", utils.ProgramName),
	Run:   runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(*cobra.Command, []string) {
	// TODO 这里处理version子命令
	fmt.Println(utils.ProgramName + ": 1.0.0")
}
