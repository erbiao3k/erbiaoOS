package initialize

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/utils/file"
	"strings"
)

// NginxMainCfg 生成nginx主配置文件
func NginxMainCfg(IPs []string) {
	upstreamConf := ""
	for _, ip := range IPs {
		upstreamConf = upstreamConf + "        server " + ip + ":6443 max_fails=3 fail_timeout=30s;\n"
	}
	cfg := strings.ReplaceAll(customConst.NginxMainConf, "upstreamConf", upstreamConf)
	file.Create(customConst.K8sMasterCfgDir+"nginx.conf", cfg)
	file.Create(customConst.K8sNodeCfgDir+"nginx.conf", cfg)
}

// NginxSystemd 生成nginx主配置文件
func NginxSystemd() {
	file.Create(customConst.K8sMasterCfgDir+"nginx.service", customConst.NginxSystemd)
	file.Create(customConst.K8sNodeCfgDir+"nginx.service", customConst.NginxSystemd)
}
