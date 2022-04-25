package kubectl

import (
	customConst "erbiaoOS/const"
	"erbiaoOS/pkg/cert"
	"erbiaoOS/utils"
	"fmt"
)

const (
	kubectlTag          = "admin"
	kubectlConfigDir    = "/root/.kube/"
	kubectlConfigfile   = kubectlConfigDir + "config"
	adminPrivateKeyFile = customConst.K8sSslDir + "admin-key.pem"
	adminPublicKeyFile  = customConst.K8sSslDir + "admin.pem"

	kubeControllerManagerTag            = "system:kube-controller-manager"
	kubeControllerManagerKubeConfig     = customConst.K8sCfgDir + "kube-controller-manager.kubeconfig"
	kubeControllerManagerPublicKeyFile  = customConst.K8sCfgDir + "kube-controller-manager.pem"
	kubeControllerManagerPrivateKeyFile = customConst.K8sCfgDir + "kube-controller-manager-key.pem"
)

var (
	// kubectl 管理客户端初始化指令
	kubectlSetClusterCmd     = fmt.Sprintf("kubectl config set-cluster kubernetes --certificate-authority=%s --embed-certs=true --server=https://%s:6443 --kubeconfig=%s\n", cert.CaPubilcKeyFile, utils.CurrentIP, kubectlConfigfile)
	kubectlSetCredentialsCmd = fmt.Sprintf("kubectl config set-credentials %s --client-certificate=%s --client-key=%s --embed-certs=true --kubeconfig=%s\n", kubectlTag, adminPublicKeyFile, adminPrivateKeyFile, kubectlConfigfile)
	kubectlSetContextCmd     = fmt.Sprintf("kubectl config set-context kubernetes --cluster=kubernetes --user=%s --kubeconfig=%s\n", kubectlTag, kubectlConfigfile)
	kubectlUseContextCmd     = fmt.Sprintf("kubectl config use-context kubernetes --kubeconfig=%s\n", kubectlConfigfile)

	controllerManagerSetClusterCmd     = fmt.Sprintf("kubectl config set-cluster kubernetes --certificate-authority=%s --embed-certs=true --server=https://%s:6443 --kubeconfig=%s\n", cert.CaPubilcKeyFile, utils.CurrentIP, kubeControllerManagerKubeConfig)
	controllerManagerSetCredentialsCmd = fmt.Sprintf("kubectl config set-credentials %s --client-certificate=%s --client-key=%s --embed-certs=true --kubeconfig=%s\n", kubeControllerManagerTag, kubeControllerManagerPublicKeyFile, kubeControllerManagerPrivateKeyFile, kubeControllerManagerKubeConfig)
	controllerManagerSetContextCmd     = fmt.Sprintf("kubectl config set-context %s --cluster=kubernetes --user=%s --kubeconfig=%s\n", kubeControllerManagerTag, kubeControllerManagerTag, kubeControllerManagerKubeConfig)
	controllerManagerUseContextCmd     = fmt.Sprintf("kubectl config use-context %s --kubeconfig=%s\n", kubeControllerManagerTag, kubectlConfigfile)
)
