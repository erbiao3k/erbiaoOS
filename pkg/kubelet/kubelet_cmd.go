package kubelet

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"strings"
)

// config 生成kubelet配置文件
func config() {
	for _, ip := range append(setting.K8sMasterIPs, setting.K8sNodeIPs...) {
		cfg := strings.ReplaceAll(cfgContent, "currentKubeletIP", ip)
		file.Create(myConst.TempDir+ip+"/kubelet", cfg)
	}
}

// systemdScript 生成kubelet配置文件以及systemd管理脚本
func systemdScript() {
	for _, hosts := range setting.K8sClusterHost {
		for _, host := range hosts {
			cfg := strings.ReplaceAll(systemd, "kubeletDataDir", host.DataDir)
			file.Create(myConst.TempDir+host.LanIp+"/kubelet.service", cfg)
		}

	}
}

func Start() {
	config()
	systemdScript()

	utils.ExecCmd(setClusterCmd)
	utils.ExecCmd(setCredentialsCmd)
	utils.ExecCmd(setContextCmd)
	utils.ExecCmd(useContextCmd)
	utils.ExecCmd(clusterrolebindingDelete)
	utils.ExecCmd(clusterrolebindingCreate)
	utils.ExecCmd(approveNode)

	for _, hosts := range setting.K8sClusterHost {
		for _, host := range hosts {
			sshd.Upload(host.LanIp, host.User, host.Password, host.Port, myConst.TempDir+host.LanIp+"/kubelet", myConst.K8sCfgDir)
			sshd.Upload(host.LanIp, host.User, host.Password, host.Port, myConst.TempDir+host.LanIp+"/kubelet.service", myConst.SystemdServiceDir)
			sshd.Upload(host.LanIp, host.User, host.Password, host.Port, kubeconfig, myConst.K8sCfgDir)
			sshd.RemoteSshExec(host.LanIp, host.User, host.Password, host.Port, restartCmd)
		}
	}
}
