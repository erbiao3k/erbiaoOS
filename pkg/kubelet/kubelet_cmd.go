package kubelet

import (
	"erbiaoOS/pkg/cert"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/num"
	sshd2 "erbiaoOS/utils/sshd"
	"erbiaoOS/vars"
	"fmt"
	"strings"
)

// config 生成kubelet配置文件
func config() {
	for _, ip := range append(vars.K8sMasterIPs, vars.K8sNodeIPs...) {
		cfg := strings.ReplaceAll(cfgContent, "currentKubeletIP", ip)
		file.Create(vars.TempDir+ip+"/kubelet", cfg)
	}
}

// systemdScript 生成kubelet配置文件以及systemd管理脚本
func systemdScript(hostInfo []vars.HostInfo) {
	for _, host := range hostInfo {
		cfg := strings.ReplaceAll(systemd, "kubeletDataDir", utils.LargeDisk(&host, "kubelet"))
		file.Create(vars.TempDir+host.LanIp+"/kubelet.service", cfg)
	}
}

func Start() {

	_, _, K8sClusterHost := vars.ClusterHostInfo()

	var (
		setClusterCmd            = fmt.Sprintf(vars.SetClusterCmd, cert.CaPubilcKeyFile, vars.EnterpointAddr(), kubeconfig)
		setCredentialsCmd        = fmt.Sprintf(vars.KubeletSetCredentialsCmd, kubeletCredentials, num.RandomString, kubeconfig)
		setContextCmd            = fmt.Sprintf(vars.SetContextCmd, context, user, kubeconfig)
		useContextCmd            = fmt.Sprintf(vars.UseContextCmd, context, kubeconfig)
		clusterrolebindingDelete = fmt.Sprintf(vars.ClusterrolebindingDelete, clusterrolebinding)
		clusterrolebindingCreate = fmt.Sprintf(vars.ClusterrolebindingCreate, clusterrolebinding, clusterrole, user)
		approveNode              = "kubectl certificate approve `kubectl get csr|awk '/node/{print $1}'`"
	)

	config()
	systemdScript(K8sClusterHost)
	cmds := []string{setClusterCmd, setCredentialsCmd, setContextCmd, useContextCmd, clusterrolebindingDelete, clusterrolebindingCreate}
	utils.MultiExecCmd(cmds)

	for _, host := range K8sClusterHost {

		sshd2.Upload(&host, vars.TempDir+host.LanIp+"/kubelet", vars.K8sCfgDir)
		sshd2.Upload(&host, vars.TempDir+host.LanIp+"/kubelet.service", vars.SystemdServiceDir)
		sshd2.Upload(&host, kubeconfig, vars.K8sCfgDir)
		sshd2.RemoteExec(&host, restartCmd)
	}
	utils.ExecCmd(approveNode)

}
