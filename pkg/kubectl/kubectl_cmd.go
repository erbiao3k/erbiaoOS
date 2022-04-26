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

	utils.ExecCmd(setClusterCmd)
	utils.ExecCmd(setCredentialsCmd)
	utils.ExecCmd(setContextCmd)
	utils.ExecCmd(useContextCmd)

	for _, hosts := range setting.K8sClusterHost {
		for _, host := range hosts {
			if host.LanIp == utils.CurrentIP {
				continue
			}
			sshd.Upload(host.LanIp, host.User, host.Password, host.Port, kubeconfig, myConst.KubectlConfigDir)
		}
	}
}
