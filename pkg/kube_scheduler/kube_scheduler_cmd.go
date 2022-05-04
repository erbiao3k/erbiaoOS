package kube_scheduler

import (
	"erbiaoOS/pkg/cert"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	sshd2 "erbiaoOS/utils/sshd"
	"erbiaoOS/vars"
	"fmt"
)

// systemdScript 生成systemd管理脚本
func systemdScript() {
	file.Create(vars.TempDir+"kube-scheduler.service", systemd)
}

// Start 初始化scheduler集群
func Start() {

	K8sMasterHost, _, _ := vars.ClusterHostInfo()

	var (
		setClusterCmd     = fmt.Sprintf(vars.SetClusterCmd, cert.CaPubilcKeyFile, vars.EnterpointAddr(), kubeconfig)
		setCredentialsCmd = fmt.Sprintf(vars.SetCredentialsCmd, user, publicKeyFile, privateKeyFile, kubeconfig)
		setContextCmd     = fmt.Sprintf(vars.SetContextCmd, context, user, kubeconfig)
		useContextCmd     = fmt.Sprintf(vars.UseContextCmd, context, kubeconfig)
	)

	systemdScript()

	cmds := []string{setClusterCmd, setCredentialsCmd, setContextCmd, useContextCmd}
	utils.MultiExecCmd(cmds)

	for _, host := range K8sMasterHost {

		sshd2.Upload(&host, vars.TempDir+"kube-scheduler.service", vars.SystemdServiceDir)
		sshd2.Upload(&host, kubeconfig, vars.K8sCfgDir)
		sshd2.RemoteExec(&host, restartCmd)
	}
}
