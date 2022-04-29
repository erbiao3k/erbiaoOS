package kubectl

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/config"
	"erbiaoOS/utils"
	"erbiaoOS/utils/login/sshd"
	"erbiaoOS/utils/net"
	"os"
)

// InitKubectl 初始化kubectl客户端
func InitKubectl() {
	os.Mkdir(myConst.KubectlConfigDir, 0600)

	cmds := []string{setClusterCmd, setCredentialsCmd, setContextCmd, useContextCmd, clusterrolebindingDelete, clusterrolebindingCreate}

	utils.MultiExecCmd(cmds)

	for _, host := range config.K8sClusterHost {
		if host.LanIp == net.CurrentIP {
			continue
		}

		sshd.Upload(&host, kubeconfig, myConst.KubectlConfigDir)
	}
}
