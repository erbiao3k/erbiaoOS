package kubelet

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/cert"
	config2 "erbiaoOS/pkg/config"
	"erbiaoOS/utils"
	"erbiaoOS/utils/file"
	"erbiaoOS/utils/login/sshd"
	"fmt"
	"strings"
)

// config 生成kubelet配置文件
func config() {
	for _, ip := range append(myConst.K8sMasterIPs, myConst.K8sNodeIPs...) {
		cfg := strings.ReplaceAll(cfgContent, "currentKubeletIP", ip)
		file.Create(myConst.TempDir+ip+"/kubelet", cfg)
	}
}

// systemdScript 生成kubelet配置文件以及systemd管理脚本
func systemdScript() {
	for _, host := range config2.K8sClusterHost {
		cfg := strings.ReplaceAll(systemd, "kubeletDataDir", host.DataDir)
		file.Create(myConst.TempDir+host.LanIp+"/kubelet.service", cfg)
	}
}

func Start() {
	var (
		setClusterCmd            = fmt.Sprintf(myConst.SetClusterCmd, cert.CaPubilcKeyFile, config2.EnterpointAddr(), kubeconfig)
		setCredentialsCmd        = fmt.Sprintf(myConst.KubeletSetCredentialsCmd, kubeletCredentials, utils.RandomString, kubeconfig)
		setContextCmd            = fmt.Sprintf(myConst.SetContextCmd, context, user, kubeconfig)
		useContextCmd            = fmt.Sprintf(myConst.UseContextCmd, context, kubeconfig)
		clusterrolebindingDelete = fmt.Sprintf(myConst.ClusterrolebindingDelete, clusterrolebinding)
		clusterrolebindingCreate = fmt.Sprintf(myConst.ClusterrolebindingCreate, clusterrolebinding, clusterrole, user)
		approveNode              = "kubectl certificate approve `kubectl get csr|awk '/node/{print $1}'`"
	)

	config()
	systemdScript()
	cmds := []string{setClusterCmd, setCredentialsCmd, setContextCmd, useContextCmd, clusterrolebindingDelete, clusterrolebindingCreate}
	utils.MultiExecCmd(cmds)

	for _, host := range config2.K8sClusterHost {

		sshd.Upload(&host, myConst.TempDir+host.LanIp+"/kubelet", myConst.K8sCfgDir)
		sshd.Upload(&host, myConst.TempDir+host.LanIp+"/kubelet.service", myConst.SystemdServiceDir)
		sshd.Upload(&host, kubeconfig, myConst.K8sCfgDir)
		sshd.RemoteExec(&host, restartCmd)
	}
	utils.ExecCmd(approveNode)

}
