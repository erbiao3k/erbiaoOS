package initialize

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/pkg/file"
	"erbiaoOS/setting"
	"strings"
)

// KubeletCfg 生成kubelet配置文件
func KubeletCfg(fileSaveDir string, hosts []string) {
	for _, ip := range hosts {
		cfg := strings.ReplaceAll(customConst.KubeletCfg, "currentKubeletIP", ip)
		file.Create(fileSaveDir+ip+"-kubelet", cfg)
	}
}

// KubeletSystemdScript 生成kubelet配置文件以及systemd管理脚本
func KubeletSystemdScript(fileSaveDir string, host []setting.HostInfo) {

	cfg := strings.ReplaceAll(customConst.KubeletSystemd, "kubeletDataDir", kubeletDataDir)
	file.Create(fileSaveDir+ip+"-kubelet.service", cfg)
}
