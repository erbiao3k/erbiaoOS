package initialize

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/pkg/file"
	"erbiaoOS/setting"
	"strings"
)

// KubeProxySystemdScript 生成kube-proxy systemd管理脚本
func KubeProxySystemdScript(dir string, hosts []setting.HostInfo) {
	for _, host := range hosts {
		cfg := strings.ReplaceAll(customConst.KubeProxySystemd, "kubeProxyDataDir", host.DataDir)
		file.Create(dir+host.LanIp+"-kube-proxy.service", cfg)
	}
}
