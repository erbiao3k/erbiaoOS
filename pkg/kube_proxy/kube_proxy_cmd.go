package kube_proxy

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"strings"
)

// config 生成配置文件
func config() {
	for _, ip := range append(setting.K8sMasterIPs, setting.K8sNodeIPs...) {
		cfg := strings.ReplaceAll(cfgContent, "currentKubeproxyIP", ip)
		file.Create(myConst.TempDir+ip+"/kube-proxy", cfg)
	}
}

// systemdScript 生成systemd管理脚本
func systemdScript() {
	for _, host := range append(setting.K8sMasterHost, setting.K8sNodeHost...) {
		cfg := strings.ReplaceAll(systemd, "kubeProxyDataDir", host.DataDir)
		file.Create(myConst.TempDir+host.LanIp+"/kube-proxy.service", cfg)
	}
}

func Start() {
	config()
	systemdScript()
	utils.ExecCmd(setClusterCmd)
	utils.ExecCmd(setCredentialsCmd)
	utils.ExecCmd(setContextCmd)
	utils.ExecCmd(useContextCmd)

	for _, host := range append(setting.K8sMasterHost, setting.K8sNodeHost...) {
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, myConst.TempDir+host.LanIp+"/kube-proxy.service", myConst.SystemdServiceDir)
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, myConst.TempDir+host.LanIp+"/kube-proxy", myConst.K8sCfgDir)
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, kubeconfig, myConst.K8sCfgDir)
		sshd.RemoteSshExec(host.LanIp, host.User, host.Password, host.Port, restartCmd)
	}
}
