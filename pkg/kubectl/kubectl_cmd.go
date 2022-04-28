package kubectl

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/login/sshd"
	"os"
)

// InitKubectl 初始化kubectl客户端
func InitKubectl() {
	os.Mkdir(myConst.KubectlConfigDir, 0600)

	cmds := []string{setClusterCmd, setCredentialsCmd, setContextCmd, useContextCmd, clusterrolebindingDelete, clusterrolebindingCreate}

	utils.MultiExecCmd(cmds)

	for _, host := range setting.K8sClusterHost {
		if host.LanIp == utils.CurrentIP {
			continue
		}
		hostInfo := &sshd.Info{
			LanIp:    host.LanIp,
			User:     host.User,
			Password: host.Password,
			Port:     host.Port,
		}

		sshd.Upload(hostInfo, kubeconfig, myConst.KubectlConfigDir)
	}
}
