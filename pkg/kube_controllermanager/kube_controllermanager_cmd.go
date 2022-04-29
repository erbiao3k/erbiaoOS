package kube_controllermanager

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/config"
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

	cmds := []string{setClusterCmd, setCredentialsCmd, setContextCmd, useContextCmd}
	utils.MultiExecCmd(cmds)

	for _, host := range config.K8sMasterHost {

		sshd.Upload(&host, myConst.TempDir+"kube-controller-manager.service", myConst.SystemdServiceDir)
		sshd.Upload(&host, kubeconfig, myConst.K8sCfgDir)
		sshd.RemoteExec(&host, restartCmd)
	}
}
