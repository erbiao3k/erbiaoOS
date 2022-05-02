package kubectl

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/cert"
	"erbiaoOS/pkg/config"
	"erbiaoOS/utils"
	"erbiaoOS/utils/login/sshd"
	"erbiaoOS/utils/net"
	"fmt"
	"os"
)

// InitKubectl 初始化kubectl客户端
func InitKubectl() {
	var (

		// kubectl 管理客户端初始化指令
		setClusterCmd            = fmt.Sprintf(myConst.SetClusterCmd, cert.CaPubilcKeyFile, config.EnterpointAddr(), kubeconfig)
		setCredentialsCmd        = fmt.Sprintf(myConst.SetCredentialsCmd, user, publicKeyFile, privateKeyFile, kubeconfig)
		setContextCmd            = fmt.Sprintf(myConst.SetContextCmd, context, user, kubeconfig)
		useContextCmd            = fmt.Sprintf(myConst.UseContextCmd, context, kubeconfig)
		clusterrolebindingDelete = fmt.Sprintf(myConst.ClusterrolebindingDelete, credentials)
		clusterrolebindingCreate = fmt.Sprintf(myConst.ClusterrolebindingCreate, clusterrolebinding, clusterrole, clusterrolebindingUser)
	)

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
