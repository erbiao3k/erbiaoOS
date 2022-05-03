package kubectl

import (
	"erbiaoOS/pkg/cert"
	"erbiaoOS/utils"
	"erbiaoOS/utils/login/sshd"
	"erbiaoOS/utils/net"
	"erbiaoOS/vars"
	"fmt"
	"os"
)

// InitKubectl 初始化kubectl客户端
func InitKubectl() {

	_, _, K8sClusterHost := vars.ClusterHostInfo()

	var (

		// kubectl 管理客户端初始化指令
		setClusterCmd            = fmt.Sprintf(vars.SetClusterCmd, cert.CaPubilcKeyFile, vars.EnterpointAddr(), kubeconfig)
		setCredentialsCmd        = fmt.Sprintf(vars.SetCredentialsCmd, user, publicKeyFile, privateKeyFile, kubeconfig)
		setContextCmd            = fmt.Sprintf(vars.SetContextCmd, context, user, kubeconfig)
		useContextCmd            = fmt.Sprintf(vars.UseContextCmd, context, kubeconfig)
		clusterrolebindingDelete = fmt.Sprintf(vars.ClusterrolebindingDelete, credentials)
		clusterrolebindingCreate = fmt.Sprintf(vars.ClusterrolebindingCreate, clusterrolebinding, clusterrole, clusterrolebindingUser)
	)

	os.Mkdir(vars.KubectlConfigDir, 0600)

	cmds := []string{setClusterCmd, setCredentialsCmd, setContextCmd, useContextCmd, clusterrolebindingDelete, clusterrolebindingCreate}

	utils.MultiExecCmd(cmds)

	for _, host := range K8sClusterHost {
		if host.LanIp == net.CurrentIP {
			continue
		}

		sshd.Upload(&host, kubeconfig, vars.KubectlConfigDir)
	}
}
