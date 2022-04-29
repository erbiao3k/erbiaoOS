package kube_proxy

import (
	myConst "erbiaoOS/const"
	config2 "erbiaoOS/pkg/config"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"strings"
)

// config 生成配置文件
func config() {
	for _, ip := range append(config2.K8sMasterIPs, config2.K8sNodeIPs...) {

		cfg := strings.ReplaceAll(cfgContent, "currentKubeproxyIP", ip)

		file.Create(myConst.TempDir+ip+"/kube-proxy", cfg)
	}
}

// systemdScript 生成systemd管理脚本
func systemdScript() {
	for _, host := range config2.K8sClusterHost {
		cfg := strings.ReplaceAll(systemd, "kubeProxyDataDir", host.DataDir)
		file.Create(myConst.TempDir+host.LanIp+"/kube-proxy.service", cfg)
	}
}

func Start() {
	config()
	systemdScript()
	cmds := []string{setClusterCmd, setCredentialsCmd, setContextCmd, useContextCmd}
	utils.MultiExecCmd(cmds)

	for _, host := range config2.K8sClusterHost {
		sshd.Upload(&host, myConst.TempDir+host.LanIp+"/kube-proxy.service", myConst.SystemdServiceDir)
		sshd.Upload(&host, myConst.TempDir+host.LanIp+"/kube-proxy", myConst.K8sCfgDir)
		sshd.Upload(&host, kubeconfig, myConst.K8sCfgDir)
		sshd.RemoteExec(&host, restartCmd)
	}
}
