package cmd

import (
	"erbiaoOS/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var rootShow = `
本工具用于二进制版本k8s集群的自动化部署配置：
	1、初始化：
		集群（开发中）
		高可用集群（开发中）
	2、节点管理：
		新增master节点（开发中）
		新增node节点（开发中）

		删除master节点（开发中）
		删除node节点（开发中）

	3、监控：
		一键配置（开发中）
	4、告警：
		一键配置（开发中）
`
var rootCmd = &cobra.Command{
	Use:   utils.ProgramName,
	Short: "本工具用于二进制版本k8s集群的自动化部署配置",
	Long:  rootShow,
	Run:   runRoot,
}

func Execute() {
	rootCmd.Execute()
}

func runRoot(*cobra.Command, []string) {
	fmt.Printf(rootShow)
}
