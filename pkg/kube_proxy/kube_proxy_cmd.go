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
	for _, host := range setting.K8sClusterHost {
		cfg := strings.ReplaceAll(systemd, "kubeProxyDataDir", host.DataDir)
		file.Create(myConst.TempDir+host.LanIp+"/kube-proxy.service", cfg)
	}
}

func Start() {
	config()
	systemdScript()
	cmds := []string{setClusterCmd, setCredentialsCmd, setContextCmd, useContextCmd}
	utils.MultiExecCmd(cmds)

	for _, host := range setting.K8sClusterHost {
		hostInfo := &sshd.Info{
			LanIp:    host.LanIp,
			User:     host.User,
			Password: host.Password,
			Port:     host.Port,
		}
		sshd.Upload(hostInfo, myConst.TempDir+host.LanIp+"/kube-proxy.service", myConst.SystemdServiceDir)
		sshd.Upload(hostInfo, myConst.TempDir+host.LanIp+"/kube-proxy", myConst.K8sCfgDir)
		sshd.Upload(hostInfo, kubeconfig, myConst.K8sCfgDir)
		sshd.RemoteSshExec(hostInfo, restartCmd)
	}
}
