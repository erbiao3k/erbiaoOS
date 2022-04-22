package kube_proxy

import (
	"erbiaoOS/setting"
	"erbiaoOS/utils/file"
	"strings"
)

// cfg 生成kube-proxy配置文件
func cfg(dir string, IPs []string) {
	for _, ip := range IPs {
		cfg := strings.ReplaceAll(cfgContent, "currentKubeproxyIP", ip)
		file.Create(dir+ip+"-kube-proxy", cfg)
	}
}

// systemdScript 生成kube-proxy systemd管理脚本
func systemdScript(dir string, hosts []setting.HostInfo) {
	for _, host := range hosts {
		cfg := strings.ReplaceAll(systemd, "kubeProxyDataDir", host.DataDir)
		file.Create(dir+host.LanIp+"-kube-proxy.service", cfg)
	}
}
