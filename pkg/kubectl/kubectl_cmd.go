package kubectl

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/login/sshd"
	"log"
	"os"
)

// InitKubectl 初始化kubectl客户端
func InitKubectl() {
	os.Mkdir(myConst.KubectlConfigDir, 0600)

	log.Println(setClusterCmd, "\n", setCredentialsCmd, "\n", setContextCmd, "\n", useContextCmd, "\n", clusterrolebindingDelete, "\n", clusterrolebindingCreate)

	utils.ExecCmd(setClusterCmd)
	utils.ExecCmd(setCredentialsCmd)
	utils.ExecCmd(setContextCmd)
	utils.ExecCmd(useContextCmd)
	utils.ExecCmd(clusterrolebindingDelete)
	utils.ExecCmd(clusterrolebindingCreate)

	for _, host := range append(setting.K8sMasterHost, setting.K8sNodeHost...) {
		if host.LanIp == utils.CurrentIP {
			continue
		}
		sshd.Upload(host.LanIp, host.User, host.Password, host.Port, kubeconfig, myConst.KubectlConfigDir)
	}
}
