package initialize

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/pkg/file"
	"strings"
)

// KubeProxyCfg 生成kube-proxy配置文件
func KubeProxyCfg(dir string, IPs []string) {
	for _, ip := range IPs {
		cfg := strings.ReplaceAll(customConst.KubeProxyCfg, "currentKubeproxyIP", ip)
		file.Create(dir+ip+"-kube-proxy", cfg)
	}
}
