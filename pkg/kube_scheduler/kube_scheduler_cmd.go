package kube_scheduler

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
)

// systemdScript 生成systemd管理脚本
func systemdScript() {
	file.Create(myConst.TempDir+"kube-scheduler.service", systemd)
}

// Start 初始化scheduler集群
func Start() {
	systemdScript()

	cmds := []string{setClusterCmd, setCredentialsCmd, setContextCmd, useContextCmd}
	utils.MultiExecCmd(cmds)

	for _, host := range setting.K8sMasterHost {
		hostInfo := &sshd.Info{
			LanIp:    host.LanIp,
			User:     host.User,
			Password: host.Password,
			Port:     host.Port,
		}

		sshd.Upload(hostInfo, myConst.TempDir+"kube-scheduler.service", myConst.SystemdServiceDir)
		sshd.Upload(hostInfo, kubeconfig, myConst.K8sCfgDir)
		sshd.RemoteSshExec(hostInfo, restartCmd)
	}
}
