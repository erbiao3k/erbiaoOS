package kube_scheduler

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/cert"
	"erbiaoOS/pkg/config"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"fmt"
)

// systemdScript 生成systemd管理脚本
func systemdScript() {
	file.Create(myConst.TempDir+"kube-scheduler.service", systemd)
}

// Start 初始化scheduler集群
func Start() {

	var (
		setClusterCmd     = fmt.Sprintf(myConst.SetClusterCmd, cert.CaPubilcKeyFile, config.EnterpointAddr(), kubeconfig)
		setCredentialsCmd = fmt.Sprintf(myConst.SetCredentialsCmd, user, publicKeyFile, privateKeyFile, kubeconfig)
		setContextCmd     = fmt.Sprintf(myConst.SetContextCmd, context, user, kubeconfig)
		useContextCmd     = fmt.Sprintf(myConst.UseContextCmd, context, kubeconfig)
	)

	systemdScript()

	cmds := []string{setClusterCmd, setCredentialsCmd, setContextCmd, useContextCmd}
	utils.MultiExecCmd(cmds)

	for _, host := range config.K8sMasterHost {

		sshd.Upload(&host, myConst.TempDir+"kube-scheduler.service", myConst.SystemdServiceDir)
		sshd.Upload(&host, kubeconfig, myConst.K8sCfgDir)
		sshd.RemoteExec(&host, restartCmd)
	}
}
