package kubectl

import (
	"bytes"
	myConst "erbiaoOS/const"
	"erbiaoOS/setting"
	"erbiaoOS/utils"
	"erbiaoOS/utils/login/sshd"
	"log"
	"os"
	"os/exec"
)

// InitKubectl 初始化kubectl客户端
func InitKubectl() {
	os.Mkdir(myConst.KubectlConfigDir, 0600)
	//
	cmd := exec.Command("bash", "-c", setClusterCmd+setCredentialsCmd+setContextCmd+useContextCmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal("初始化kubectl客户端异常：", err)
	}

	for _, hosts := range setting.K8sServer {
		for _, host := range hosts {
			if host.LanIp == utils.CurrentIP {
				continue
			}
			sshd.Upload(host.LanIp, host.User, host.Password, host.Port, kubeconfig, myConst.KubectlConfigDir)
		}
	}
}
