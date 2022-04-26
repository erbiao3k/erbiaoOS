package kubectl

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/cert"
	"erbiaoOS/setting"
	"fmt"
)

const (
	user    = "admin"
	context = "kubernetes"

	kubeconfig     = myConst.KubectlConfigDir + "config"
	privateKeyFile = myConst.K8sSslDir + "admin-key.pem"
	publicKeyFile  = myConst.K8sSslDir + "admin.pem"
)

var (

	// kubectl 管理客户端初始化指令
	setClusterCmd     = fmt.Sprintf(myConst.SetClusterCmd, cert.CaPubilcKeyFile, setting.RandMasterIP, kubeconfig)
	setCredentialsCmd = fmt.Sprintf(myConst.SetCredentialsCmd, user, publicKeyFile, privateKeyFile, kubeconfig)
	setContextCmd     = fmt.Sprintf(myConst.SetContextCmd, context, user, kubeconfig)
	useContextCmd     = fmt.Sprintf(myConst.UseContextCmd, context, kubeconfig)
)
