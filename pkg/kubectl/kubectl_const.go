package kubectl

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/pkg/cert"
	"erbiaoOS/utils"
	"fmt"
)

const (
	kubectlUser         = "admin"
	kubectlContext      = "kubernetes"
	kubectlConfigDir    = "/root/.kube/"
	kubectlConfigfile   = kubectlConfigDir + "config"
	adminPrivateKeyFile = customConst.K8sSslDir + "admin-key.pem"
	adminPublicKeyFile  = customConst.K8sSslDir + "admin.pem"

	kubeControllerManagerUser           = "system:kube-controller-manager"
	kubeControllerManagerContext        = kubeControllerManagerUser
	kubeControllerManagerKubeConfig     = customConst.K8sCfgDir + "kube-controller-manager.kubeconfig"
	kubeControllerManagerPublicKeyFile  = customConst.K8sCfgDir + "kube-controller-manager.pem"
	kubeControllerManagerPrivateKeyFile = customConst.K8sCfgDir + "kube-controller-manager-key.pem"

	kubeSchedulerUser           = "system:kube-scheduler"
	kubeSchedulerContext        = kubeSchedulerUser
	kubeSchedulerKubeConfig     = customConst.K8sCfgDir + "kube-scheduler.kubeconfig"
	kubeSchedulerPublicKeyFile  = customConst.K8sCfgDir + "kube-scheduler.pem"
	kubeSchedulerPrivateKeyFile = customConst.K8sCfgDir + "kube-scheduler-key.pem"
)

var (
	// kubectl 管理客户端初始化指令
	kubectlSetClusterCmd     = fmt.Sprintf(customConst.SetClusterCmd, cert.CaPubilcKeyFile, utils.CurrentIP, kubectlConfigfile)
	kubectlSetCredentialsCmd = fmt.Sprintf(customConst.SetCredentialsCmd, kubectlUser, adminPublicKeyFile, adminPrivateKeyFile, kubectlConfigfile)
	kubectlSetContextCmd     = fmt.Sprintf(customConst.SetContextCmd, kubectlContext, kubectlUser, kubectlConfigfile)
	kubectlUseContextCmd     = fmt.Sprintf(customConst.UseContextCmd, kubectlContext, kubectlConfigfile)

	controllerManagerSetClusterCmd     = fmt.Sprintf(customConst.SetClusterCmd, cert.CaPubilcKeyFile, utils.CurrentIP, kubeControllerManagerKubeConfig)
	controllerManagerSetCredentialsCmd = fmt.Sprintf(customConst.SetCredentialsCmd, kubeControllerManagerUser, kubeControllerManagerPublicKeyFile, kubeControllerManagerPrivateKeyFile, kubeControllerManagerKubeConfig)
	controllerManagerSetContextCmd     = fmt.Sprintf(customConst.SetContextCmd, kubeControllerManagerContext, kubeControllerManagerUser, kubeControllerManagerKubeConfig)
	controllerManagerUseContextCmd     = fmt.Sprintf(customConst.UseContextCmd, kubeControllerManagerContext, kubeControllerManagerKubeConfig)

	kubeSchedulerSetClusterCmd     = fmt.Sprintf(customConst.SetClusterCmd, cert.CaPubilcKeyFile, utils.CurrentIP, kubeSchedulerKubeConfig)
	kubeSchedulerSetCredentialsCmd = fmt.Sprintf(customConst.SetCredentialsCmd, kubeSchedulerUser, kubeSchedulerPublicKeyFile, kubeSchedulerPrivateKeyFile, kubeSchedulerKubeConfig)
	kubeSchedulerSetContextCmd     = fmt.Sprintf(customConst.SetContextCmd, kubeSchedulerContext, kubeSchedulerUser, kubeSchedulerKubeConfig)
	kubeSchedulerUseContextCmd     = fmt.Sprintf(customConst.UseContextCmd, kubeSchedulerContext, kubeSchedulerKubeConfig)
)
