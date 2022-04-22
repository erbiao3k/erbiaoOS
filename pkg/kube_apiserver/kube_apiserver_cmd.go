package kube_apiserver

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/utils/file"
	"strconv"
	"strings"
)

// systemdScript 生成kube-apiserver systemd管理脚本
func systemdScript(masterIPs []string, etcdServerUrls string) {
	apiserverCount := len(masterIPs)
	for _, ip := range masterIPs {
		cfg := strings.ReplaceAll(systemd, "currentIPaddr", ip)
		cfg = strings.ReplaceAll(cfg, "etcdServerUrls", etcdServerUrls)
		cfg = strings.ReplaceAll(cfg, "apiserverCount", strconv.Itoa(apiserverCount))
		file.Create(customConst.K8sMasterCfgDir+ip+"-kube-apiserver.service", cfg)
	}

}
