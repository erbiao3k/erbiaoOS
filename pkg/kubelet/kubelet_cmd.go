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
	for _, host := range setting.K8sClusterHost {
		cfg := strings.ReplaceAll(systemd, "kubeletDataDir", host.DataDir)
		file.Create(myConst.TempDir+host.LanIp+"/kubelet.service", cfg)
	}
}

func Start() {
	config()
	systemdScript()
	cmds := []string{setClusterCmd, setCredentialsCmd, setContextCmd, useContextCmd, clusterrolebindingDelete, clusterrolebindingCreate}
	utils.MultiExecCmd(cmds)

	for _, host := range setting.K8sClusterHost {

		sshd.Upload(&host, myConst.TempDir+host.LanIp+"/kubelet", myConst.K8sCfgDir)
		sshd.Upload(&host, myConst.TempDir+host.LanIp+"/kubelet.service", myConst.SystemdServiceDir)
		sshd.Upload(&host, kubeconfig, myConst.K8sCfgDir)
		sshd.RemoteExec(&host, restartCmd)
	}
	utils.ExecCmd(approveNode)

}
