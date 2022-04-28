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

		sshd.Upload(&host, kubeconfig, myConst.KubectlConfigDir)
	}
}
