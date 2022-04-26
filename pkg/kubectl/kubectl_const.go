package kubectl

import (
	myConst "erbiaoOS/const"
	"erbiaoOS/pkg/cert"
	"erbiaoOS/setting"
	"fmt"
)

const (
	kubectlUser    = "admin"
	kubectlContext = "kubernetes"

	kubectlConfigfile   = myConst.KubectlConfigDir + "config"
	adminPrivateKeyFile = myConst.K8sSslDir + "admin-key.pem"
	adminPublicKeyFile  = myConst.K8sSslDir + "admin.pem"
)

var (

	// kubectl 管理客户端初始化指令
	kubectlSetClusterCmd     = fmt.Sprintf(myConst.SetClusterCmd, cert.CaPubilcKeyFile, setting.RandMasterIP, kubectlConfigfile)
	kubectlSetCredentialsCmd = fmt.Sprintf(myConst.SetCredentialsCmd, kubectlUser, adminPublicKeyFile, adminPrivateKeyFile, kubectlConfigfile)
	kubectlSetContextCmd     = fmt.Sprintf(myConst.SetContextCmd, kubectlContext, kubectlUser, kubectlConfigfile)
	kubectlUseContextCmd     = fmt.Sprintf(myConst.UseContextCmd, kubectlContext, kubectlConfigfile)
)
