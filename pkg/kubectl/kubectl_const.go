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

	kubeSchedulerUser           = "system:kube-scheduler"
	kubeSchedulerContext        = kubeSchedulerUser
	kubeSchedulerKubeConfig     = myConst.TempDir + "kube-scheduler.kubeconfig"
	kubeSchedulerPublicKeyFile  = myConst.K8sSslDir + "kube-scheduler.pem"
	kubeSchedulerPrivateKeyFile = myConst.K8sSslDir + "kube-scheduler-key.pem"
)

var (

	// kubectl 管理客户端初始化指令
	kubectlSetClusterCmd     = fmt.Sprintf(myConst.SetClusterCmd, cert.CaPubilcKeyFile, setting.RandMasterIP, kubectlConfigfile)
	kubectlSetCredentialsCmd = fmt.Sprintf(myConst.SetCredentialsCmd, kubectlUser, adminPublicKeyFile, adminPrivateKeyFile, kubectlConfigfile)
	kubectlSetContextCmd     = fmt.Sprintf(myConst.SetContextCmd, kubectlContext, kubectlUser, kubectlConfigfile)
	kubectlUseContextCmd     = fmt.Sprintf(myConst.UseContextCmd, kubectlContext, kubectlConfigfile)

	kubeSchedulerSetClusterCmd     = fmt.Sprintf(myConst.SetClusterCmd, cert.CaPubilcKeyFile, setting.RandMasterIP, kubeSchedulerKubeConfig)
	kubeSchedulerSetCredentialsCmd = fmt.Sprintf(myConst.SetCredentialsCmd, kubeSchedulerUser, kubeSchedulerPublicKeyFile, kubeSchedulerPrivateKeyFile, kubeSchedulerKubeConfig)
	kubeSchedulerSetContextCmd     = fmt.Sprintf(myConst.SetContextCmd, kubeSchedulerContext, kubeSchedulerUser, kubeSchedulerKubeConfig)
	kubeSchedulerUseContextCmd     = fmt.Sprintf(myConst.UseContextCmd, kubeSchedulerContext, kubeSchedulerKubeConfig)
)
