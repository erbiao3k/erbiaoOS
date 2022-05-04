package kube_proxy

import (
	"erbiaoOS/pkg/cert"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	sshd2 "erbiaoOS/utils/sshd"
	"erbiaoOS/vars"
	"fmt"
	"strings"
)

// config 生成配置文件
func config() {
	for _, ip := range append(vars.K8sMasterIPs, vars.K8sNodeIPs...) {

		cfg := strings.ReplaceAll(cfgContent, "currentKubeproxyIP", ip)

		file.Create(vars.TempDir+ip+"/kube-proxy", cfg)
	}
}

// systemdScript 生成systemd管理脚本
func systemdScript(hostInfo []vars.HostInfo) {
	for _, host := range hostInfo {
		cfg := strings.ReplaceAll(systemd, "kubeProxyDataDir", utils.LargeDisk(&host, "kube-proxy"))
		file.Create(vars.TempDir+host.LanIp+"/kube-proxy.service", cfg)
	}
}

func Start() {
	_, _, K8sClusterHost := vars.ClusterHostInfo()

	var (
		setClusterCmd     = fmt.Sprintf(vars.SetClusterCmd, cert.CaPubilcKeyFile, vars.EnterpointAddr(), kubeconfig)
		setCredentialsCmd = fmt.Sprintf(vars.SetCredentialsCmd, user, publicKeyFile, privateKeyFile, kubeconfig)
		setContextCmd     = fmt.Sprintf(vars.SetContextCmd, context, user, kubeconfig)
		useContextCmd     = fmt.Sprintf(vars.UseContextCmd, context, kubeconfig)
	)

	config()
	systemdScript(K8sClusterHost)
	cmds := []string{setClusterCmd, setCredentialsCmd, setContextCmd, useContextCmd}
	utils.MultiExecCmd(cmds)

	for _, host := range K8sClusterHost {
		sshd2.Upload(&host, vars.TempDir+host.LanIp+"/kube-proxy.service", vars.SystemdServiceDir)
		sshd2.Upload(&host, vars.TempDir+host.LanIp+"/kube-proxy", vars.K8sCfgDir)
		sshd2.Upload(&host, kubeconfig, vars.K8sCfgDir)
		sshd2.RemoteExec(&host, restartCmd)
	}
}
