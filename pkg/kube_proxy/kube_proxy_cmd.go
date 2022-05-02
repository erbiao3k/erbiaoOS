package kube_proxy

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/cert"
	config2 "erbiaoOS/pkg/config"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"fmt"
	"strings"
)

// config 生成配置文件
func config() {
	for _, ip := range append(myConst.K8sMasterIPs, myConst.K8sNodeIPs...) {

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

	var (
		setClusterCmd     = fmt.Sprintf(myConst.SetClusterCmd, cert.CaPubilcKeyFile, config2.EnterpointAddr(), kubeconfig)
		setCredentialsCmd = fmt.Sprintf(myConst.SetCredentialsCmd, user, publicKeyFile, privateKeyFile, kubeconfig)
		setContextCmd     = fmt.Sprintf(myConst.SetContextCmd, context, user, kubeconfig)
		useContextCmd     = fmt.Sprintf(myConst.UseContextCmd, context, kubeconfig)
	)

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
