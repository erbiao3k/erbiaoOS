package nginx

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/utils/file"
	"strings"
)

// MainCfg 生成nginx主配置文件
func MainCfg(IPs []string) {
	upstreamConf := ""
	for _, ip := range IPs {
		upstreamConf = upstreamConf + "        server " + ip + ":6443 max_fails=3 fail_timeout=30s;\n"
	}
	cfg := strings.ReplaceAll(mainConf, "upstreamConf", upstreamConf)
	file.Create(customConst.K8sMasterCfgDir+"nginx.conf", cfg)
	file.Create(customConst.K8sNodeCfgDir+"nginx.conf", cfg)
}

// systemdScript 生成nginx主配置文件
func systemdScript() {
	file.Create(customConst.K8sMasterCfgDir+"nginx.service", systemd)
	file.Create(customConst.K8sNodeCfgDir+"nginx.service", systemd)
}
