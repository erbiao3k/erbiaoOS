package kube_controllermanager

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
)

// systemdScript 生成kube-controller-manager systemd管理脚本
func systemdScript() {
	file.Create(myConst.TempDir+"kube-controller-manager.service", systemd)
}

func Start() {
	systemdScript()
	utils.ExecCmd(setClusterCmd)
	utils.ExecCmd(setCredentialsCmd)
	utils.ExecCmd(setContextCmd)
	utils.ExecCmd(useContextCmd)

	for _, host := range setting.K8sMasterHost {
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, myConst.TempDir+"kube-controller-manager.service", myConst.SystemdServiceDir)
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, kubeconfig, myConst.K8sCfgDir)
		sshd.RemoteSshExec(host.LanIp, host.User, host.Password, host.Port, restartCmd)
	}
}
