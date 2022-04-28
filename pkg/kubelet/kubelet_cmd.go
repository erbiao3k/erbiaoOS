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

		hostInfo := &sshd.Info{
			LanIp:    host.LanIp,
			User:     host.User,
			Password: host.Password,
			Port:     host.Port,
		}

		sshd.Upload(hostInfo, myConst.TempDir+host.LanIp+"/kubelet", myConst.K8sCfgDir)
		sshd.Upload(hostInfo, myConst.TempDir+host.LanIp+"/kubelet.service", myConst.SystemdServiceDir)
		sshd.Upload(hostInfo, kubeconfig, myConst.K8sCfgDir)
		sshd.RemoteSshExec(hostInfo, restartCmd)
	}
	utils.ExecCmd(approveNode)

}
