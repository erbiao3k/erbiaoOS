package kube_controllermanager

import (
	"erbiaoOS/pkg/cert"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"erbiaoOS/vars"
	"fmt"
)

// systemdScript 生成kube-controller-manager systemd管理脚本
func systemdScript() {
	file.Create(vars.TempDir+"kube-controller-manager.service", systemd)
}

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

		sshd.Upload(&host, vars.TempDir+"kube-controller-manager.service", vars.SystemdServiceDir)
		sshd.Upload(&host, kubeconfig, vars.K8sCfgDir)
		sshd.RemoteExec(&host, restartCmd)
	}
}
