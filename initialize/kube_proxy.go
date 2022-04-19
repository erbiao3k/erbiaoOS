package initialize

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/pkg/file"
	"erbiaoOS/setting"
	"strings"
)

// KubeProxyCfg 生成kube-proxy配置文件
func KubeProxyCfg(dir string, IPs []string) {
	for _, ip := range IPs {
		cfg := strings.ReplaceAll(customConst.KubeProxyCfg, "currentKubeproxyIP", ip)
		file.Create(dir+ip+"-kube-proxy", cfg)
	}
}

// KubeProxySystemdScript 生成kube-proxy systemd管理脚本
func KubeProxySystemdScript(dir string, hosts []setting.HostInfo) {
	for _, host := range hosts {
		cfg := strings.ReplaceAll(customConst.KubeProxySystemd, "kubeProxyDataDir", host.DataDir)
		file.Create(dir+host.LanIp+"-kube-proxy.service", cfg)
	}
}
