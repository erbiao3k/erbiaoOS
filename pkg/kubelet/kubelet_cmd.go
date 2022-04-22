package kubelet

import (
	"erbiaoOS/setting"
	"erbiaoOS/utils/file"
	"strings"
)

// cfg 生成kubelet配置文件
func cfg(fileSaveDir string, IPs []string) {
	for _, ip := range IPs {
		cfg := strings.ReplaceAll(cfgContent, "currentKubeletIP", ip)
		file.Create(fileSaveDir+ip+"-kubelet", cfg)
	}
}

// systemdScript 生成kubelet配置文件以及systemd管理脚本
func systemdScript(fileSaveDir string, hosts []setting.HostInfo) {
	for _, host := range hosts {
		cfg := strings.ReplaceAll(systemd, "kubeletDataDir", host.DataDir)
		file.Create(fileSaveDir+host.LanIp+"-kubelet.service", cfg)
	}
}
